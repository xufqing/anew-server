package asset

import (
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"regexp"
	"strings"
	"sync"
	"time"
)

// 参考 https://github.com/dejavuzhou/felix
const (
	wsMsgCmd    = "cmd"
	wsMsgResize = "resizePty"
)

type wsWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex //互斥锁
}

func (w *wsWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

// 封装ssh session
type SSHSession struct {
	// calling Write() to write data into ssh server
	StdinPipe io.WriteCloser
	// Write() be called to receive data from ssh server
	Output  *wsWriter
	Session *ssh.Session
}

type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

func WriteByteMessage(w *wsWriter, wsConn *websocket.Conn) error {
	if w.buffer.Len() != 0 {
		err := wsConn.WriteMessage(websocket.TextMessage, w.buffer.Bytes())
		if err != nil {
			return err
		}
		w.buffer.Reset()
	}
	return nil
}

func NewSSHSession(cols, rows int, sshClient *ssh.Client) (*SSHSession, error) {
	sshSession, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	stdinP, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}

	writer := new(wsWriter)
	//ssh.stdout and stderr will write output into comboWriter
	sshSession.Stdout = writer
	sshSession.Stderr = writer

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		return nil, err
	}
	// Start remote shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &SSHSession{StdinPipe: stdinP, Output: writer, Session: sshSession}, nil
}

func (s *SSHSession) Close() {
	if s.Session != nil {
		s.Session.Close()
	}

}

//利用正则表达式压缩字符串，去除空格或制表符
func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

//ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (s *SSHSession) receiveWsMsg(wsConn *websocket.Conn, exitCh chan bool, key string) {
	//tells other go routine quit
	defer setQuit(exitCh)
	var cmdStr string
	for {
		select {
		case <-exitCh:
			return
		default:
			//read websocket msg
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				code := err.(*websocket.CloseError).Code
				if code == 1000 || code == 1001 {
					hub.delete(key)
				}
				common.Log.Debugf("reading webSocket message failed\t%s", err)
				return
			}
			//unmashal bytes into struct
			var msgObj wsMsg
			err = json.Unmarshal(wsData, &msgObj)
			switch msgObj.Type {
			case wsMsgResize:
				if msgObj.Cols > 0 && msgObj.Rows > 0 {
					if err := s.Session.WindowChange(msgObj.Rows, msgObj.Cols); err != nil {
						common.Log.Debugf("ssh pty change windows size failed:\t", err)
					}
				}
			case wsMsgCmd:
				// 写入cmd到shell stdin
				if _, err := s.StdinPipe.Write(utils.Str2Bytes(msgObj.Cmd)); err != nil {
					common.Log.Debugf("ws cmd bytes write to ssh.stdin pipe failed:\t", err)
				}
				if msgObj.Cmd == "\r" || msgObj.Cmd == "\n" {
					if cmdStr != "" {
						fmt.Println(compressStr(cmdStr))
						cmdStr = ""
					}
				} else {
					//matched,_ :=regexp.MatchString("[\\u0001-\\u0003]",msgObj.Cmd)
					//if matched{
					//	cmdStr =cmdStr + msgObj.Cmd
					//} else{
					//	fmt.Println("特殊符号")
					//}
					switch msgObj.Cmd {
					// ctrl + c
					case "\u0003":
						cmdStr = ""
					// 退格
					case "\u007F":
						lastStr := cmdStr[len(cmdStr)-1:]
						cmdStr = strings.TrimSuffix(cmdStr, lastStr)
					default:
						cmdStr = cmdStr + msgObj.Cmd
					}

				}
			}
		}
	}
}

func (s *SSHSession) sendOutput(wsConn *websocket.Conn, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)
	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(120))
	//for range time.Tick(120 * time.Millisecond){}
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			//write combine output bytes into websocket response
			if err := WriteByteMessage(s.Output, wsConn); err != nil {
				common.Log.Debugf("ssh sending combo output to webSocket failed:\t", err)
				return
			}
		case <-exitCh:
			return
		}
	}
}

func (s *SSHSession) sessionWait(quitChan chan bool, key string) {
	if err := s.Session.Wait(); err != nil {
		common.Log.Debugf("ssh session wait failed:\t%s", err)
		hub.delete(key)
		setQuit(quitChan)
	}
}

func setQuit(ch chan bool) {
	ch <- true
}
