package connection

import (
	"WebSocketIM/datasource"
	"WebSocketIM/delayqueue"
	nodeClient "WebSocketIM/grpc/node/client"
	pb "WebSocketIM/grpc/node/proto"
	"WebSocketIM/mq"
	static "WebSocketIM/static"
	"fmt"
	"github.com/wonderivan/logger"
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
	message := data.(static.Message)
	fmt.Printf("消息过期：%+v\n", message)
	switch message.MsgType {
	case SEND:
		{
			//此message存在于ack队列中
			readAt, exist := ackMap[message.MsgId]
			//并且在expire前得到了ack
			if exist {
				// 按照已读消息入库
				_, err := datasource.Insert(datasource.InsertMessageRead, message.MsgId, message.MsgType, message.Data, message.FromUid, message.ToUid, message.CreateAt, readAt)
				if err != nil {
					logger.Error("message %+v insert failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
				}
				// 移除ackMap中该条记录
				delete(ackMap, message.MsgId)
			} else {
				// 没有得到ack 按照未读消息入库
				_, err := datasource.Insert(datasource.InsertMessage, message.MsgId, message.MsgType, message.Data, message.FromUid, message.ToUid, message.CreateAt)
				if err != nil {
					logger.Error("message %+v insert failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
					return
				}
			}
		}
	case READ_ACK:
		{
			//此message存在于ack队列中
			readAt, exist := ackMap[message.MsgId]
			if exist {
				//发现ack记录 说明该条记录已经按照未读入库了  直接修改
				_, err := datasource.Update(datasource.UpdateMessageStatusToRead, readAt, message.MsgId)
				if err != nil {
					logger.Error("message %+v read failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
				}
				//移除该条记录
				delete(ackMap, message.MsgId)
			}
			// 未发现ack记录  说明该消息已经在两秒内发现了ack 并按照已读入库了  这里不用做任何处理
		}
	case REVOKE:
		{
			//修改message状态
			_, err := datasource.Update(datasource.UpdateMessageStatusToRevoke, message.RevokeAt, message.MsgId)
			if err != nil {
				logger.Error("message %+v revoke failed at %s.", message.MsgId, time.Now().Format("2006/01/02 15:04:05"))
			}
		}
	}

}
func InitConsumer() (err error) {
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
		msg := mq.MyClient.GetPayLoad(Channel).(static.Message)
		switch msg.MsgType {
		case SEND:
			{
				logger.Info("get message is %+v\n", msg)

				connId, exist := UtoC[msg.ToUid]
				if exist {
					// 目标用户在线 将该条send消息存入延时队列 为的是接受在几秒钟内到达的readACK 之后再写入数据库 减小写库压力
					dq.TPush(msg)
					c := AllConnection[connId]
					c.WriteMessage(msg)
				} else {
					// 目标用户可能在其他节点上
					targetNodeName, targetNodeAddr, err := nodeClient.FindUser(msg.ToUid)
					if err == nil {
						// 成功查询到目标用户所在的node 发送消息
						convertMsg := ConvertMessage2(msg)
						success := nodeClient.SendMessage(targetNodeName, targetNodeAddr, &convertMsg)
						if success {
							goto flag
						}
					}
					//目标用户不在线 直接入库 如果数据库写入压力太大 可以做一个缓冲队列 定时批量写入
					_, err = datasource.Insert(datasource.InsertMessage, msg.MsgId, msg.MsgType, msg.Data, msg.FromUid, msg.ToUid, msg.CreateAt)
					if err != nil {
						logger.Error("message %+v insert failed at %s.", msg.MsgId, time.Now().Format("2006/01/02 15:04:05"))
					}
				}
			flag:
				//响应发送方
				connId, exist = UtoC[msg.FromUid]
				if exist {
					c := AllConnection[connId]
					c.WriteMessage(static.Message{
						MsgType: SENT,
						MsgId:   msg.MsgId,
					})
				} else {
					// 目标用户可能在其他节点上
					targetNodeName, targetNodeAddr, err := nodeClient.FindUser(msg.FromUid)
					if err == nil {
						// 成功查询到目标用户所在的node 发送消息
						convertMsg := ConvertMessage2(msg)
						nodeClient.SendMessage(targetNodeName, targetNodeAddr, &convertMsg)
					}
				}
			}
		case READ_ACK:
			{
				logger.Info("get ack is %+v\n", msg)
				ackMap[msg.MsgId] = msg.ReadAt
				dq.TPush(msg)
				// 如果接收方在线 响应接受方
				connId, exist := UtoC[msg.ToUid]
				if exist {
					c := AllConnection[connId]
					c.WriteMessage(static.Message{
						MsgType: READ_ACK,
						MsgId:   msg.MsgId,
					})
				} else {
					// 目标用户可能在其他节点上
					targetNodeName, targetNodeAddr, err := nodeClient.FindUser(msg.ToUid)
					if err == nil {
						// 成功查询到目标用户所在的node 发送消息
						convertMsg := ConvertMessage2(msg)
						nodeClient.SendMessage(targetNodeName, targetNodeAddr, &convertMsg)
					}
					// 目标用户不在线 直接修改
				}
			}
		case REVOKE:
			{
				logger.Info("get revoke is %+v\n", msg)
				dq.TPush(msg)
				// 如果接收方在线 响应接受方
				connId, exist := UtoC[msg.ToUid]
				if exist {
					c := AllConnection[connId]
					c.WriteMessage(static.Message{
						MsgType:  REVOKE,
						MsgId:    msg.MsgId,
						RevokeAt: msg.RevokeAt,
					})
				} else {
					// 目标用户可能在其他节点上
					targetNodeName, targetNodeAddr, err := nodeClient.FindUser(msg.ToUid)
					if err == nil {
						// 成功查询到目标用户所在的node 发送消息
						convertMsg := ConvertMessage2(msg)
						nodeClient.SendMessage(targetNodeName, targetNodeAddr, &convertMsg)
					}
				}
				//响应发起方
				connId, exist = UtoC[msg.FromUid]
				if exist {
					c := AllConnection[connId]
					c.WriteMessage(static.Message{
						MsgType:  REVOKE_ACK,
						MsgId:    msg.MsgId,
						RevokeAt: msg.RevokeAt,
					})
				} else {
					// 目标用户可能在其他节点上
					targetNodeName, targetNodeAddr, err := nodeClient.FindUser(msg.FromUid)
					if err == nil {
						// 成功查询到目标用户所在的node 发送消息
						convertMsg := ConvertMessage2(msg)
						nodeClient.SendMessage(targetNodeName, targetNodeAddr, &convertMsg)
					}
				}
			}

		}

	}
}

func ConvertMessage2(msg static.Message) pb.SendMessageRequest_Message {
	retMsg := pb.SendMessageRequest_Message{}
	retMsg.MsgId = msg.MsgId
	retMsg.MsgType = int32(msg.MsgType)
	retMsg.Data = msg.Data
	retMsg.FromUid = msg.FromUid
	retMsg.ToUid = msg.ToUid
	retMsg.CreateAt = msg.CreateAt.Unix()
	retMsg.IsRead = int32(msg.IsRead)
	retMsg.ReadAt = msg.ReadAt.Unix()
	retMsg.IsRevoke = int32(msg.IsRevoke)
	retMsg.RevokeAt = msg.RevokeAt.Unix()

	return retMsg
}
