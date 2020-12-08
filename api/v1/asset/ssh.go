package asset

import (
	system2 "anew-server/api/v1/system"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/sshx"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	// 心跳间隔
	heartBeatPeriod = 10 * time.Second
	// 心跳最大重试次数
	HeartBeatMaxRetryCount = 3
)

var (
	UpGrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 全局变量，管理ssh连接
	hub ConnectionHub
)

// 连接仓库，管理所有连接ssh的客户端
type ConnectionHub struct {
	// 客户端集合(用户id为每个socket key)
	Clients map[string]*Connection
}

type Connection struct {
	// 当前socket key
	Key string
	// d当前连接
	Conn *websocket.Conn
	// 当前登录用户名
	UserName string
	// 当前登录用户姓名
	Name string
	// 主机名称
	HostName string
	// 主机名称
	User string
	// 主机地址
	IpAddress string
	// 主机端口
	Port string
	// 接入时间
	ConnectionTime models.LocalTime
	// 上次活跃时间
	LastActiveTime models.LocalTime
	// 重试次数
	RetryCount uint
}

type Writer struct {
	b  bytes.Buffer
	mu sync.Mutex // 互斥锁
}

func (w *Writer) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.b.Write(p)
}

func (w *Writer) Read() ([]byte, int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	p := w.b.Bytes()
	buf := make([]byte, len(p))
	read, err := w.b.Read(buf)
	w.b.Reset()
	if err != nil {
		return nil, 0, err
	}
	return buf, read, err
}

// 启动消息中心仓库
func StartConnectionHub() {
	// 初始化仓库
	hub.Clients = make(map[string]*Connection)
}

func (h *ConnectionHub) append(key string, client *Connection) {
	h.Clients[key] = client
}

func (h *ConnectionHub) get(key string) (*Connection, error) {
	var err error
	client := h.Clients[key]
	if client == nil {
		err = errors.New("连接不存在")
		return client, err
	}
	return client, err

}

func (h *ConnectionHub) delete(key string) {
	delete(h.Clients, key)
}

// websocket
func SSHTunnel(c *gin.Context) {
	// 绑定参数,获取ssh连接的信息
	var req request.SSHTunnelReq
	errq := c.Bind(&req)
	if errq != nil {
		common.Log.Error("参数绑定失败")
		return
	}
	s := service.New()
	host, err := s.GetHostById(req.HostId)
	width, _ := strconv.Atoi(c.DefaultQuery("width", "1024"))
	height, _ := strconv.Atoi(c.DefaultQuery("height", "768"))
	if err != nil {
		common.Log.Error(err.Error())
		return
	}

	// 获取当前登录用户
	user := system2.GetCurrentUserFromCache(c)
	websocketKey := c.Request.Header.Get("Sec-WebSocket-Key")
	client, err := hub.get(websocketKey)
	if err != nil {
		ws, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			common.Log.Error("创建消息连接失败", err)
			return
		}
		// 注册到消息仓库
		client = &Connection{
			Key:       websocketKey,
			Conn:      ws,
			UserName:  user.(system.SysUser).Username,
			Name:      user.(system.SysUser).Name,
			HostName:  host.HostName,
			User:      host.User,
			IpAddress: host.IpAddress,
			Port:      host.Port,
			ConnectionTime: models.LocalTime{
				Time: time.Now(),
			},
			LastActiveTime: models.LocalTime{
				Time: time.Now(),
			},
		}
		// 加入连接仓库
		hub.append(websocketKey, client)
	}
	conf, err := sshx.NewAuthConfig(host.User, host.Password, "", "")
	if err != nil {
		common.Log.Error(err.Error())
		return
	}
	// 默认密码模式
	if host.AuthType == "key" {
		conf, err = sshx.NewAuthConfig(host.User, "", host.PrivateKey, host.KeyPassphrase)
		if err != nil {
			common.Log.Error(err.Error())
			return
		}
	}
	addr := fmt.Sprintf("%s:%s", host.IpAddress, host.Port)
	sshClient, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		common.Log.Error(err.Error())
	}
	session, err := sshClient.NewSession()
	if err != nil {
		common.Log.Error(err.Error())
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	// Pty模式
	if err := session.RequestPty("xterm", width, height, modes); err != nil {
		common.Log.Error(err.Error())
	}

	//
	var b Writer
	session.Stdout = &b
	session.Stderr = &b

	stdinPipe, err := session.StdinPipe()

	if err != nil {
		common.Log.Error(err.Error())
	}

	if err := session.Shell(); err != nil {
		common.Log.Error(err.Error())
	}

	go client.heartBeat()
	go client.send(&b)
	client.receive(stdinPipe)
}

func (c *Connection) receive(b io.WriteCloser) {
	for {
		_, message, err := c.Conn.ReadMessage()

		// 记录活跃时间
		c.LastActiveTime = models.LocalTime{
			Time: time.Now(),
		}
		if err != nil {
			common.Log.Debugf("接收数据失败,(%s:%s) 可能已断开:\t%s", c.IpAddress, c.Key, err)
			hub.delete(c.Key)
			c.RetryCount = 100 //不要再发心跳包了
			break
		}
		// 接收前端心跳，重置心跳计数器
		if string(message) == "ssh-heart-beat-in" {
			c.LastActiveTime = models.LocalTime{
				Time: time.Now(),
			}
			c.RetryCount = 0
		} else {
			_, err = b.Write(message)
			if err != nil {
				common.Log.Debugf("写入数据失败,(%s:%s) 可能已断开:\t%s", c.IpAddress, c.Key, err)
			}
		}
	}
}

func (c *Connection) send(b *Writer) {
	for {
		p, n, err := b.Read()
		if err != nil {
			continue
		}
		if n > 0 {
			WriteByteMessage(c.Conn, p)
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}
func (c *Connection) heartBeat() {
	// 创建定时器, 超出指定时间间隔, 向前端发送ping消息心跳
	ticker := time.NewTicker(heartBeatPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
		if err := recover(); err != nil {
			common.Log.Debugf("发送心跳失败,(%s:%s) 可能已断开:\t%s", c.IpAddress, c.Key, err)
		}
	}()
loop:
	for {
		select {
		// 到心跳检测时间
		case <-ticker.C:
			//common.Log.Debug("当前活跃连接数:", len(hub.Clients))
			last := time.Now().Sub(c.LastActiveTime.Time)
			if c.RetryCount > HeartBeatMaxRetryCount {
				common.Log.Debugf("多次心跳无响应,(%s:%s) 可能已断开", c.IpAddress, c.Key)
				hub.delete(c.Key)
				break loop
			}
			if last > heartBeatPeriod {
				// 发送心跳
				WriteMessage(c.Conn, "ssh-heart-beat-out")
				c.RetryCount++
			}
		}
	}
}
func WriteMessage(ws *websocket.Conn, message string) {
	WriteByteMessage(ws, []byte(message))
}

func WriteByteMessage(ws *websocket.Conn, p []byte) {
	err := ws.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		common.Log.Debug("Ws写入消息失败:", err)
	}
}

// 获取连接列表
func GetConnections(c *gin.Context) {
	var resp []response.ConnectionResp
	for client, _ := range hub.Clients {
		var connStruct response.ConnectionResp
		connStruct.Key = hub.Clients[client].Key
		connStruct.UserName = hub.Clients[client].UserName
		connStruct.Name = hub.Clients[client].Name
		connStruct.HostName = hub.Clients[client].HostName
		connStruct.IpAddress = hub.Clients[client].IpAddress
		connStruct.Port = hub.Clients[client].Port
		connStruct.ConnectionTime = hub.Clients[client].ConnectionTime
		connStruct.LastActiveTime = hub.Clients[client].LastActiveTime
		resp = append(resp, connStruct)
	}

	response.SuccessWithData(resp)
}

// 注销已登录的连接
func DeleteConnectionByKey(c *gin.Context) {
	var req request.KeyReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	conn, err := hub.get(req.Key)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 关闭socket
	conn.Conn.Close()
	// 连接池删除对象
	hub.delete(req.Key)
	response.Success()
}
