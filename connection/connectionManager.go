package connection

import (
	nodeClient "WebSocketIM/grpc/node/client"
	"WebSocketIM/static"
	"errors"
	"sync"
)

const (
	MaxConnNum = 10000
)

//用户map key为当前连接的connId value为对应Connection 通过维护connId到userId的映射来存储登陆关系
var AllConnection = make(map[int64]*Connection, MaxConnNum)

//connId到userId的映射
var CtoU = make(map[int64]string, MaxConnNum)

//userId到connId的映射
var UtoC = make(map[string]int64, MaxConnNum)

// 最大的连接ID，每次连接都加1 处理
var MaxConnId int64

var addLock sync.Mutex

var idLock sync.Mutex

func GetConn(connId int64) (conn *Connection, err error) {
	conn, exist := AllConnection[connId]
	if !exist {
		err = errors.New("connId不存在")
	}
	return
}

func AddConn(conn *Connection) (err error) {
	addLock.Lock()
	defer addLock.Unlock()
	_, exist := AllConnection[conn.ConnId]
	if exist {
		err = errors.New("connId已存在")
	} else {
		AllConnection[conn.ConnId] = conn
	}
	return
}

func RemoveConn(connId int64) (err error) {
	_, exist := AllConnection[connId]
	if exist {
		err = errors.New("connId不存在")
	} else {
		delete(AllConnection, connId)
	}
	return
}

func Login(userId string, conn *Connection) (err error) {
	if userId == "" {
		err = errors.New("userId为空")
	} else {
		conn.UserId = userId
		// todo 多端登陆逻辑
		// 暂时只实现一端登陆
		// 强制下线
		if oldConnId, exist := UtoC[userId]; exist {
			//移除connID → UserId 的映射
			delete(CtoU, oldConnId)
			forceToLogout(oldConnId)
		}

		// 记录新的id映射关系
		UtoC[userId] = conn.ConnId
		CtoU[conn.ConnId] = userId

		// 通告zookeeper 忽略失败的远程调用
		nodeClient.UserCheckIn(userId)
	}
	return
}

func Logout(userId string) {
	// todo 这里还应该关闭连接
	connId := UtoC[userId]
	delete(UtoC, userId)
	delete(CtoU, connId)
	conn, exist := AllConnection[connId]
	// 释放该连接
	if exist {
		conn.Close()
	}
	// 通告zookeeper 忽略失败的远程调用
	nodeClient.UserCheckOut(userId)
}

//账号被他人登录 强制下线
func forceToLogout(connId int64) {
	if connection, exist := AllConnection[connId]; exist {
		//todo 向被强制下线的客户端发送一条通知
		//sendMessage("你被强制下线啦")
		delete(AllConnection, connId)
		connection.Close()
	}
}

// 获取一个新的连接id
func GetNewConnId() (ret int64) {
	idLock.Lock()
	MaxConnId++
	ret = MaxConnId
	idLock.Unlock()
	return
}

// 查询某一个用户是否在线
func IsOnline(userId string) bool {
	_, exist := UtoC[userId]
	return exist
}

// 供grpc接口使用 从参数中获取msg 写入指定userId的conn中
func SendMessage(userId string, msg static.Message) error {
	connId, exist := UtoC[userId]
	if !exist {
		return errors.New("userId不存在 目标用户不在本节点")
	}
	conn, exist := AllConnection[connId]
	if !exist {
		return errors.New("connId不存在 节点数据异常")
	}

	select {
	case conn.inChan <- msg:
		return nil
	case <-conn.closeChan: // closeChan 感知 conn断开
		return errors.New("本节点与目标用户连接中断 发送失败")
	}

}
