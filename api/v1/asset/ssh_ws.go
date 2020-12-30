package asset

import (
	system2 "anew-server/api/v1/system"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/sshx"
	"anew-server/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net/http"
	"strconv"
	"time"
)

var (
	UpGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
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
	// 客户端集合
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
	// 主机用户名
	User string
	// 主机地址
	IpAddress string
	// 主机端口
	Port string
	// 接入时间
	ConnectTime models.LocalTime
	// ssh session 对象
	SshClient *ssh.Client
	// sftp session 对象
	SFTPClient *sftp.Client
	// 录像文件名称
	CastFileName string
}

func (c *Connection) close() {
	if c.Conn != nil {
		c.Conn.Close()
	}
	if c.SshClient != nil {
		c.SshClient.Close()
	}
	if c.SFTPClient != nil {
		c.SshClient.Close()
	}
	if c.Key != "" {
		hub.delete(c.Key)
	}
}

// 启动Ssh连接仓库
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
		return nil, err
	}
	return client, err

}

func (h *ConnectionHub) delete(key string) {
	delete(h.Clients, key)
}

// websocket
func SshTunnel(c *gin.Context) {
	hostId := utils.Str2Uint(c.Query("host_id"))
	s := service.New()
	host, err := s.GetHostById(hostId)
	if err != nil {
		common.Log.Error(err.Error())
		return
	}
	cols, _ := strconv.Atoi(c.DefaultQuery("cols", "120"))
	rows, _ := strconv.Atoi(c.DefaultQuery("rows", "32"))

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
			ConnectTime: models.LocalTime{
				Time: time.Now(),
			},
		}
		// 加入连接仓库
		hub.append(websocketKey, client)
	}
	// 关闭
	defer client.close()
	// 发送websocketKey给前端

	client.Conn.WriteMessage(websocket.TextMessage, utils.Str2Bytes("Anew-Sec-WebSocket-Key:"+websocketKey))
	var conf *ssh.ClientConfig
	switch host.AuthType {
	case "key":
		conf, err = sshx.NewAuthConfig(host.User, "", host.PrivateKey, host.KeyPassphrase)
		if err != nil {
			common.Log.Error(err.Error())
			return
		}
	default:
		// 默认密码模式
		conf, err = sshx.NewAuthConfig(host.User, host.Password, "", "")
		if err != nil {
			common.Log.Error(err.Error())
			return
		}
	}
	addr := fmt.Sprintf("%s:%s", host.IpAddress, host.Port)
	sshClient, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		client.Conn.WriteMessage(websocket.TextMessage, utils.Str2Bytes(fmt.Sprintf("Ssh无法连接: %s", err.Error())))
		common.Log.Errorf("Ssh无法连接: %s", err.Error())
		return
	}
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		common.Log.Errorf("SFTP无法连接: %s", err.Error())
		return
	}
	client.SshClient = sshClient
	client.SFTPClient = sftpClient
	// 创建ssh session
	sshSession, err := NewSshSession(cols, rows, sshClient)
	if err != nil {
		client.Conn.WriteMessage(websocket.TextMessage, utils.Str2Bytes(fmt.Sprintf("Ssh Session创建失败: %s", err.Error())))
		common.Log.Error("Ssh Session创建失败: %s", err.Error())
		return
	}
	defer sshSession.Session.Close()
	quitChan := make(chan bool, 3)
	go sshSession.SendOutput(client, quitChan)  // 协程输出ws和进行录像
	go sshSession.SessionWait(quitChan)
	sshSession.ReceiveWsMsg(client, quitChan)

}
