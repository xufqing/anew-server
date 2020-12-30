package asset

import (
	"anew-server/models"
	"anew-server/models/asset"
	"anew-server/pkg/asciicast2"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"strconv"
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
type SshSession struct {
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
	err := wsConn.WriteMessage(websocket.TextMessage, w.buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func NewSshSession(cols, rows int, sshClient *ssh.Client) (*SshSession, error) {
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
	return &SshSession{StdinPipe: stdinP, StdOutput: writer, Session: sshSession}, nil
}

func (s *SshSession) Close(c *Connection) {
	s.saveRecord(c)
	if s.Session != nil {
		s.Session.Close()
	}
}

//ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (s *SshSession) ReceiveWsMsg(c *Connection, exitCh chan bool) {
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
					s.Close(c)
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

func (s *SshSession) SendOutput(c *Connection, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)
	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(120))
	//for range time.Tick(120 * time.Millisecond){}
	defer tick.Stop()
	// ssh 录像
	meta := asciicast2.CastV2Header{
		Width:     130,
		Height:    30,
		Timestamp: time.Now().Unix(),
		Title:     c.Key,
		Env: &map[string]string{
			"SHELL": "/bin/bash", "TERM": "xterm-256color",
		},
	}
	startTime := time.Now()
	castFile := c.IpAddress + "_"+ strconv.FormatInt(time.Now().UnixNano(), 10) + ".cast"
	c.CastFileName = castFile
	if !utils.FileExist(common.Conf.Ssh.RecordDir) {
		_ = os.Mkdir(common.Conf.Ssh.RecordDir, 644)
	}
	cast, f := asciicast2.NewCastV2(meta, castFile)
	defer f.Close()
	for {
		select {
		case <-tick.C:
			if s.StdOutput.buffer.Len() != 0 {
				// 发送ws和创建审计录像
				if err := WriteByteMessage(s.StdOutput, c.Conn); err != nil {
					common.Log.Debugf("ssh sending combo output to webSocket failed:\t", err)
					return
				}
				// 录像
				cast.Record(startTime, s.StdOutput.buffer.Bytes())
				s.StdOutput.buffer.Reset() // 录像操作完毕后reset buffer
			}
		case <-exitCh:
			return
		}
	}
}

func (s *SshSession) SessionWait(quitChan chan bool) {
	if err := s.Session.Wait(); err != nil {
		common.Log.Debugf("ssh session wait failed:\t%s", err)
		setQuit(quitChan)
	}
}
func (s *SshSession) saveRecord(c *Connection) {
	record := asset.SshRecord{
		Key:         c.Key,
		UserName:    c.UserName,
		HostName:    c.HostName,
		IpAddress:   c.IpAddress,
		Port:        c.Port,
		User:        c.User,
		ConnectTime: c.ConnectTime,
		LogoutTime: models.LocalTime{
			Time: time.Now(),
		},
		CastFileName: c.CastFileName,
	}
	common.Mysql.Create(&record)
}

func setQuit(ch chan bool) {
	ch <- true
}
