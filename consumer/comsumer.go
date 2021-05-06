package consumer

import (
	"WebSocketIM/connection"
	"WebSocketIM/datasource"
	"WebSocketIM/delayqueue"
	"WebSocketIM/mq"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	Channel <-chan interface{}
	dq      *delayqueue.Queue
	ackMap  map[string]time.Time
	ackLock sync.Mutex
)

func myHandler(data interface{}) {
	message := data.(connection.Message)
	fmt.Printf("消息过期：%+v\n", message)
	switch message.MsgType {
	case connection.SEND:
		{
			//此message存在于ack队列中
			readAt, exist := ackMap[message.MsgId]
			//并且在expire前得到了ack
			if exist {
				// 按照已读消息入库
				_, err := datasource.Insert(datasource.InsertMessageRead, message.MsgId, message.MsgType, message.Data, message.FromUid, message.ToUid, message.CreateAt, readAt)
				if err != nil {
					log.Printf("message %+v insert failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
				}
				// 移除ackMap中该条记录
				delete(ackMap, message.MsgId)
			} else {
				// 没有得到ack 按照未读消息入库
				_, err := datasource.Insert(datasource.InsertMessage, message.MsgId, message.MsgType, message.Data, message.FromUid, message.ToUid, message.CreateAt)
				if err != nil {
					log.Printf("message %+v insert failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
				}
			}
		}
	case connection.READ_ACK:
		{
			//此message存在于ack队列中
			readAt, exist := ackMap[message.MsgId]
			if exist {
				//发现ack记录 说明该条记录已经按照未读入库了  直接修改
				_, err := datasource.Update(datasource.UpdateMessageStatusToRead, readAt, message.MsgId)
				if err != nil {
					log.Printf("message %+v read failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
				}
				//移除该条记录
				delete(ackMap, message.MsgId)
				//响应这条ack的接收方
				//如果接收方在线 响应接受方
				connId, exist := connection.UtoC[message.ToUid]
				if exist {
					c := connection.AllConnection[connId]
					c.WriteMessage(connection.Message{
						MsgType: connection.READ_ACK,
						MsgId:   message.MsgId,
						ReadAt:  readAt,
					})
				}
			}
			//如果没有记录 说明dq已经将它按照已读写入 不用处理
		}
	case connection.REVOKE:
		{
			//修改message状态
			_, err := datasource.Update(datasource.UpdateMessageStatusToRevoke, message.RevokeAt, message.MsgId)
			if err != nil {
				log.Printf("message %+v revoke failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
			}
		}
	}

}
func Init() (err error) {
	Channel, err = mq.MyClient.Subscribe(mq.Topic)
	if err != nil {
		fmt.Println("subscribe failed")
	}
	dq = delayqueue.TimeQueue(2*time.Second, 500, myHandler)
	dq.StartTimeSpying()
	ackMap = make(map[string]time.Time, 500)
	go workLoop()
	return
}

func workLoop() {
	for {
		msg := mq.MyClient.GetPayLoad(Channel).(connection.Message)
		switch msg.MsgType {
		case connection.SEND:
			{
				log.Printf("get message is %+v\n", msg)

				connId, exist := connection.UtoC[msg.ToUid]
				if exist {
					// 目标用户在线 将该条send消息存入延时队列 为的是接受在几秒钟内到达的readACK 之后再写入数据库 减小写库压力
					dq.TPush(msg)
					c := connection.AllConnection[connId]
					c.WriteMessage(msg)
				} else {
					//目标用户不在线 直接入库 如果数据库写入压力太大 可以做一个缓冲队列 定时批量写入
					_, err := datasource.Insert(datasource.InsertMessage, msg.MsgId, msg.MsgType, msg.Data, msg.FromUid, msg.ToUid, msg.CreateAt)
					if err != nil {
						log.Printf("message %+v insert failed at %s.", msg.MsgId, time.Now().Format("2006/01/02 15:04:05"))
					}
				}
				//响应发送方
				connId, exist = connection.UtoC[msg.FromUid]
				if exist {
					c := connection.AllConnection[connId]
					c.WriteMessage(connection.Message{
						MsgType: connection.SENT,
						MsgId:   msg.MsgId,
					})
				}

			}
		case connection.READ_ACK:
			{
				log.Printf("get ack is %+v\n", msg)
				ackMap[msg.MsgId] = msg.ReadAt
				dq.TPush(msg)
				// 如果接收方在线 响应接受方
				connId, exist := connection.UtoC[msg.ToUid]
				if exist {
					c := connection.AllConnection[connId]
					c.WriteMessage(connection.Message{
						MsgType: connection.READ_ACK,
						MsgId:   msg.MsgId,
					})
				}
			}
		case connection.REVOKE:
			{
				log.Printf("get revoke is %+v\n", msg)
				dq.TPush(msg)
				// 如果接收方在线 响应接受方
				connId, exist := connection.UtoC[msg.ToUid]
				if exist {
					c := connection.AllConnection[connId]
					c.WriteMessage(connection.Message{
						MsgType:  connection.REVOKE,
						MsgId:    msg.MsgId,
						RevokeAt: msg.RevokeAt,
					})
				}
				//响应发起方
				connId, exist = connection.UtoC[msg.FromUid]
				if exist {
					c := connection.AllConnection[connId]
					c.WriteMessage(connection.Message{
						MsgType:  connection.REVOKE_ACK,
						MsgId:    msg.MsgId,
						RevokeAt: msg.RevokeAt,
					})
				}
			}

		}

	}

}
