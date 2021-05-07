package static

import (
	"encoding/json"
	"time"
)

// 用来存放一些全局的变量 为的是解决多个包之间为了引用变量产生的import成环问题

var (
	Mode             = "" // 当前运行模式
	ZooKeeperAddress = "" //  zookeeper的gRPC服务地址
	Name             = "" // 当前节点名
	NodeAddress      = "" // 当前节点ip:port
	HttpAddress      = "" // 当前节点的http服务地址
)

// 客户端读写消息
type Message struct {
	// websocket.TextMessage 消息类型
	id       int       `db:"id"`
	MsgId    string    `json:"msgId" db:"msgId"`
	MsgType  int       `json:"msgType"  db:"msgType"`
	Data     string    `json:"data"  db:"data"`
	FromUid  string    `json:"fromUid"  db:"fromUid"`
	ToUid    string    `json:"toUid" db:"toUid"`
	CreateAt time.Time `json:"createAt" db:"createAt"`
	IsRead   int       `json:"isRead" db:"isRead"`
	ReadAt   time.Time `json:"readAt" db:"readAt"`
	IsRevoke int       `json:"isRevoke" db:"isRevoke"`
	RevokeAt time.Time `json:"revokeAt" db:"revokeAt"`
}

func (d Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	return json.Marshal(&struct {
		Alias
		CreateAt string `json:"createAt"`
	}{
		Alias:    Alias(d),
		CreateAt: time.Time(d.CreateAt).Format("2006-01-02 15:04:05"),
	})
}
