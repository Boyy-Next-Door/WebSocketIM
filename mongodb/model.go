package mongodb

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Message struct {
	Id      bson.ObjectId `bson:"_id"`
	MsgId   string        `bson:"msgId"`
	Type    string        `bson:"type"`
	FromUid string        `bson:"fromUid"`
	ToUid   string        `bson:"toUid"`
	IsRead  bool          `bson:"isRead"`
	Content string        `bson:"content"`
	Date    time.Time     `bson:"date"`
}

const (
	Database   = "IMdb"
	Collection = "TestModel"
)
