package websocketx

import (
	"anew-server/pkg/common"
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

//封装Connection

type Connection struct {
	wsConn    *websocket.Conn //长连接
	inChan    chan []byte     //接收通道
	outChan   chan []byte     //发出通道
	closeChan chan byte       //关闭通道
	mutex     sync.Mutex      //为了保证线程安全需要加锁
	isClosed  bool            //关闭状态
}

//封装websocket长连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	//go conn.procLoop()
	//启动读协程
	go conn.readLoop()
	//启动写协程
	go conn.writeLoop()
	return
}

//封装ReadAPI
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

//封装WriteAPI
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

//封装CloseAPI
func (conn *Connection) Close() {
	//wsConn.Close() 是线程安全的，可重入的Close
	conn.wsConn.Close()

	//一个chan只能关闭一次。保证这行代码只执行一次,加锁保证是线程安全的
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

//内部实现readAPI
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		//一直不停的读消息
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲的位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: //closeChan被关闭
			goto ERR
		}
	}
ERR:
	conn.Close()
}

//内部实现writeAPI
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func (wsConn *Connection) procLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.WriteMessage([]byte("heartbeat from server")); err != nil {
				common.Log.Debug("heartbeat fail")
				wsConn.Close()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.ReadMessage()
		if err != nil {
			common.Log.Debug("read fail")
			break
		}
		common.Log.Debug(string(msg))
		err = wsConn.WriteMessage(msg)
		if err != nil {
			common.Log.Debug("write fail")
			break
		}
	}
}