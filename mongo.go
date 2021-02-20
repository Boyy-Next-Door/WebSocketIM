package main

import (
	"WebSocketIM/mongodb"
	"WebSocketIM/mongodb/demo"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"time"
)

func main2() {
	//// 设置客户端连接配置
	//clientOptions := options.Client().ApplyURI("mongodb://159.75.5.163:27017")
	//
	//// 连接到MongoDB
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 检查连接
	//err = client.Ping(context.TODO(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Connected to MongoDB!")
	//insert one document
	data := &demo.Data{
		Id:      bson.NewObjectId(),
		Title:   "博客的标题 1",
		Des:     "博客描述信息 1",
		Content: "博客的具体内容 1",
		Date:    time.Now(),
	}

	err := mongodb.Insert(demo.Database, demo.Collection, data)
	if err != nil {
		fmt.Println("insert one doc", err)
	} else {
		fmt.Println("insert success")
	}

	//// find one with all fields
	//var result mongodb.Data
	//err = mongodb.FindOne(mongodb.Database, mongodb.Collection, bson.M{"_id": bson.ObjectIdHex("5b3db2334d661ff46ee14b9c")}, nil, &result)
	//fmt.Println("find one with all fields", result)
}
