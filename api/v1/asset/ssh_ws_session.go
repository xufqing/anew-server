package asset

import (
	"anew-server/pkg/asciicast2"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
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
	StdOutput *wsWriter
	Session   *ssh.Session
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
	return &SSHSession{StdinPipe: stdinP, StdOutput: writer, Session: sshSession}, nil
}

func (s *SSHSession) Close() {
	if s.Session != nil {
		s.Session.Close()
	}
}

//ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (s *SSHSession) ReceiveWsMsg(c *Connection, exitCh chan bool) {
	//tells other go routine quit
	for {
		select {
		case <-exitCh:
			return
		default:
			//read websocket msg
			_, wsData, err := c.Conn.ReadMessage()
			if err != nil {
				code := err.(*websocket.CloseError).Code
				if code == 1000 || code == 1001 {
					s.Close()
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
			}
		}
	}
}

func (s *SSHSession) SendOutput(c *Connection, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)
	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(120))
	//for range time.Tick(120 * time.Millisecond){}
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			// 发送ws
			if err := WriteByteMessage(s.StdOutput, c.Conn); err != nil {
				common.Log.Debugf("ssh sending combo output to webSocket failed:\t", err)
				return
			}
		case <-exitCh:
			return
		}
	}
}

func (s *SSHSession) SessionWait(quitChan chan bool) {
	if err := s.Session.Wait(); err != nil {
		common.Log.Debugf("ssh session wait failed:\t%s", err)
		setQuit(quitChan)
	}
}
func (s *SSHSession) save(buff *bytes.Buffer) {
	fmt.Println(buff.String())
}
func (s *SSHSession) GenAsciicastFile(w uint, h uint, title string, quitChan chan bool) {
	// ssh 录像
	var buff bytes.Buffer
	meta := asciicast2.CastV2Header{
		Width:     w,
		Height:    h,
		Timestamp: time.Now().Unix(),
		Title:     title,
		Env: &map[string]string{
			"SHELL": "/bin/bash", "TERM": "xterm-256color",
		},
	}
	cast := asciicast2.NewCastV2(meta, &buff)
	startTime := time.Now()

	defer setQuit(quitChan)
	defer s.save(&buff)
	tick := time.NewTicker(time.Millisecond * time.Duration(120))
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if s.StdOutput.buffer.Len() != 0 {
				cast.PushFrame(startTime, s.StdOutput.buffer.Bytes())
				//s.StdOutput.buffer.Reset()
			}
		case <-quitChan:
			return
		}
	}
}

func setQuit(ch chan bool) {
	ch <- true
}
