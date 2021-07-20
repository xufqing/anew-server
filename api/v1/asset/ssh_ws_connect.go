package asset

import (
	system2 "anew-server/api/v1/system"
	"anew-server/dao"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"anew-server/pkg/ws_session"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"sync"
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
)

type streamMap struct {
	sync.RWMutex
	innerMap map[string]*ws_session.WebsocketStream
}

var (
	SteamMap = NewStreamMap() // 全局 websocket 连接流的 map, 管理所有连接SSH的客户端
)

func NewStreamMap() *streamMap {
	sm := new(streamMap)
	sm.innerMap = make(map[string]*ws_session.WebsocketStream)
	return sm
}

func (sm *streamMap) Set(key string, value *ws_session.WebsocketStream) {
	sm.Lock()
	sm.innerMap[key] = value
	sm.Unlock()
}

func (sm *streamMap) Get(key string) (*ws_session.WebsocketStream, error) {
	sm.RLock()
	value := sm.innerMap[key]
	sm.RUnlock()
	if value == nil {
		err := errors.New("对象不存在")
		return nil, err
	}
	return value, nil
}

func (sm *streamMap) Remove(key string) {
	sm.RLock()
	delete(sm.innerMap, key)
	sm.RUnlock()
	return
}

// 连接 WebSocket
func Connect(c *gin.Context) {
	hostId := utils.Str2Uint(c.Query("host_id"))
	s := dao.New()
	host, err := s.GetHostById(hostId)
	if err != nil {
		common.Log.Error(err.Error())
		return
	}
	cols, _ := strconv.Atoi(c.DefaultQuery("cols", "120"))
	rows, _ := strconv.Atoi(c.DefaultQuery("rows", "32"))

	// 获取当前登录用户
	user := system2.GetCurrentUserFromCache(c)
	uid := uuid.NewV4().String()

	//获取SSH配置
	terminalConfig := ws_session.Config{
		User:          host.User,
		IpAddress:     host.IpAddress,
		Port:          host.Port,
		Password:      host.Password,
		PrivateKey:    host.PrivateKey,
		KeyPassphrase: host.KeyPassphrase,
		Width:         cols,
		Height:        rows,
	}

	// 获取ws连接
	ws, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		common.Log.Error("创建消息连接失败", err)
		return
	}
	terminal, err := ws_session.NewTerminal(terminalConfig)

	if err != nil {
		_ = ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
		_ = ws.Close()
		return
	}
	wsConn := ws_session.NewWsConn(ws)
	stream := ws_session.NewWebSocketSteam(terminal, wsConn, ws_session.Meta{
		TERM:      terminal.TERM,
		Width:     terminalConfig.Width,
		Height:    terminalConfig.Height,
		ConnectId: uid,
		UserName:  user.(system.SysUser).Username,
		HostName:  host.HostName,
		HostId:    host.Id,
	})

	err = stream.Terminal.Connect(stream, stream, stream)

	if err != nil {
		_ = ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
		_ = ws.Close()
		return
	}
	// 断开ws和ssh的操作
	stream.Terminal.SetCloseHandler(func() error {
		// 记录用户的操作
		if err := stream.Write2Log(); err != nil {
			return err
		}
		SteamMap.Remove(uid)
		return stream.Conn.Ws.Close()
	})

	stream.Conn.Ws.SetCloseHandler(func(code int, text string) error {
		SteamMap.Remove(uid)
		return terminal.Close()
	})

	SteamMap.Set(uid, stream)
	// 发送websocketKey给前端
	stream.Conn.WriteMessage(websocket.TextMessage, utils.Str2Bytes("Anew-Sec-WebSocket-Key:"+uid+"\r\n"))

	go func() {
		for {
			// 每5秒
			timer := time.NewTimer(5 * time.Second)
			<-timer.C

			if stream.Terminal.IsClosed() {
				_ = timer.Stop()
				break
			}
			// 如果有 10 分钟没有数据流动，则断开连接
			if time.Now().Unix()-stream.UpdatedAt.Unix() > 60*10 {
				stream.Conn.WriteMessage(websocket.TextMessage, utils.Str2Bytes("检测到终端闲置，已断开连接...\r\n"))
				_ = stream.Conn.WriteMessage(websocket.BinaryMessage, utils.Str2Bytes("检测到终端闲置，已断开连接..."))
				_ = stream.Conn.Ws.Close()
				_ = timer.Stop()
				break
			}
		}
	}()

}
