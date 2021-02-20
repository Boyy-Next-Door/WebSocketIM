package connection

import (
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
		//todo 多端登陆逻辑
		//暂时只实现一端登陆
		//强制下线
		if oldConnId, exist := UtoC[userId]; exist {
			//移除connID → UserId 的映射
			delete(CtoU, oldConnId)
			forceToLogout(oldConnId)
		}

		//记录新的id映射关系
		UtoC[userId] = conn.ConnId
		CtoU[conn.ConnId] = userId
	}
	return
}

func Logout(userId string) {
	connId := UtoC[userId]
	delete(UtoC, userId)
	delete(CtoU, connId)
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

func GetNewConnId() (ret int64) {
	idLock.Lock()
	MaxConnId++
	ret = MaxConnId
	idLock.Unlock()
	return
}
