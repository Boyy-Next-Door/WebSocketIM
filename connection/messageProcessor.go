package connection

import (
	"WebSocketIM/mq"
	static "WebSocketIM/static"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/wonderivan/logger"
	"log"
	"sync"
)

//MsgType
const (
	SEND       = 1
	REVOKE     = 2
	LOGIN      = 3
	LOGOUT     = 4
	READ_ACK   = 5
	SENT       = 6
	REVOKE_ACK = 7
	ERROR      = 8
)

type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan static.Message
	outChan   chan static.Message
	closeChan chan byte

	mutex    sync.Mutex // 对closeChan关闭上锁
	isClosed bool       // 防止closeChan被关闭多次
	ConnId   int64      // 连接Id
	UserId   string     // 该连接绑定的用户id
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan static.Message, 1000),
		outChan:   make(chan static.Message, 1000),
		closeChan: make(chan byte, 1),
		ConnId:    GetNewConnId(),
	}
	//加入连接中心
	if err = AddConn(conn); err != nil {
		conn = nil
		return
	}

	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	//启动消息处理协程
	go conn.processLoop()

	return
}

func (conn *Connection) ReadMessage() (msg static.Message, err error) {

	select {
	case msg = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
		// 这里需要从用户列表中删除该连接 否则会导致后续的消息无法正确收发
		RemoveConn(conn.ConnId)
	}
	return
}

func (conn *Connection) WriteMessage(msg static.Message) (err error) {

	select {
	case conn.outChan <- msg:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
		RemoveConn(conn.ConnId)
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
	RemoveConn(conn.ConnId)
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
		//msg.CreateAt = time.Now()
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
		data static.Message
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
		//log.Println("接收到消息", (string)(msg.Data))
		logger.Info(msg)
		//处理消息
		switch msg.MsgType {
		case SEND:
			mq.MyClient.Publish(mq.Topic, msg)
		case REVOKE:
			mq.MyClient.Publish(mq.Topic, msg)
		case LOGIN:
			//fromUid即登录id
			err := Login(msg.FromUid, conn)
			//响应客户端
			if err != nil {
				conn.WriteMessage(static.Message{
					MsgId:   msg.MsgId,
					MsgType: ERROR,
					Data:    "登陆失败: " + err.Error(),
				})
			} else {
				conn.WriteMessage(static.Message{
					MsgId:   msg.MsgId,
					MsgType: LOGIN,
					Data:    "登陆成功",
				})
			}
		case LOGOUT:
			Logout(msg.FromUid)
		case READ_ACK:
			mq.MyClient.Publish(mq.Topic, msg)
		}

		//err = wsConn.wsWrite(msg.messageType, msg.data)
		//if err != nil {
		//	log.Println("发送消息给客户端出现错误", err.Error())
		//	break
		//}
	}
}

func parseMessage(data []byte) (msg static.Message) {
	//解析数据
	json.Unmarshal(data, &msg)
	return
}
