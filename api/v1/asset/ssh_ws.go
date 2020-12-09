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
	"anew-server/pkg/utils"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"net/http"
	"strconv"
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
	// 关闭ws
	defer client.Conn.Close()

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
	sshSession, err := NewSSHSession(cols, rows, sshClient)
	defer sshSession.Close()
	quitChan := make(chan bool, 3)
	var buff = new(bytes.Buffer)
	go sshSession.sendOutput(client.Conn, quitChan)
	go sshSession.sessionWait(quitChan, client.Key)
	sshSession.receiveWsMsg(client.Conn, buff, quitChan, client.Key)
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
	// 连接池删除对象
	hub.delete(req.Key)
	// 关闭socket
	conn.Conn.Close()
	response.Success()
}
