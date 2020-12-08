package sshws


import (
"errors"
"fmt"
"github.com/gorilla/websocket"
"sync"
"time"
)

// 客户端读写消息
type WsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type Connection struct {
	wsSocket  *websocket.Conn // 底层websocket
	inChan    chan *WsMessage // 读队列
	outChan   chan *WsMessage // 写队列
	mutex     sync.Mutex      // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

func New(wsSocket *websocket.Conn, inChan chan *WsMessage, outChan chan *WsMessage, closeChan chan byte, isClosed bool) *Connection {
	return &Connection{wsSocket: wsSocket, inChan: inChan, outChan: outChan, closeChan: closeChan, isClosed: isClosed}
}

func (wsConn *Connection) ReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &WsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.Close()
closed:
}

func (wsConn *Connection) WriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.Close()
closed:
}

func (wsConn *Connection) ProcLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.Write(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.Close()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.Read()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		fmt.Println(string(msg.data))
		err = wsConn.Write(msg.messageType, msg.data)
		if err != nil {
			fmt.Println("write fail")
			break
		}
	}
}

func (wsConn *Connection) Write(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &WsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *Connection) Read() (*WsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *Connection) Close() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}
