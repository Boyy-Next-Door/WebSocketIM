package connection

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan Message
	outChan   chan Message
	closeChan chan byte

	mutex    sync.Mutex // 对closeChan关闭上锁
	isClosed bool       // 防止closeChan被关闭多次
	ConnId   int64      // 连接Id
	UserId   string     // 该连接绑定的用户id
}

// 客户端读写消息
type Message struct {
	// websocket.TextMessage 消息类型
	MsgId       string    `json:"msgId"`
	MessageType int       `json:"messageType"`
	Data        string    `json:"data"`
	FromUid     string    `json:"fromUid"`
	ToUid       string    `json:"toUid"`
	Arrived     bool      `json:"arrived"`
	CreateAt    time.Time `json:"createAt"`
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan Message, 1000),
		outChan:   make(chan Message, 1000),
		closeChan: make(chan byte, 1),
		ConnId:    GetNewConnId(),
	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	//启动消息处理协程
	go conn.processLoop()
	return
}

func (conn *Connection) ReadMessage() (msg Message, err error) {

	select {
	case msg = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) WriteMessage(msg Message) (err error) {

	select {
	case conn.outChan <- msg:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		//解析数据 封装成Message
		msg := parseMessage(data)

		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- msg:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}

	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data Message
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		//将message转换成json进行传输
		marshal, _ := json.Marshal(data)
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, marshal); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

// 处理队列中的消息
func (conn *Connection) processLoop() {
	// 处理消息队列中的消息
	// 获取到消息队列中的消息，处理完成后，发送消息给客户端
	for {
		msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("获取消息出现错误", err.Error())
			break
		}
		log.Println("接收到消息", (string)(msg.Data))
		// 修改以下内容把客户端传递的消息传递给处理程序
		//err = wsConn.wsWrite(msg.messageType, msg.data)
		if err != nil {
			log.Println("发送消息给客户端出现错误", err.Error())
			break
		}
	}
}

func parseMessage(data []byte) (msg Message) {
	//解析数据
	json.Unmarshal(data, &msg)
	return
}
