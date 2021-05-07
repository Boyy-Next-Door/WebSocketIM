package client

import (
	nodePB "WebSocketIM/grpc/node/proto"
	zkPB "WebSocketIM/grpc/zookeeper/proto"
	"WebSocketIM/static"
	"errors"
	"github.com/wonderivan/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"strings"
	"time"
)

var (
	conn     *grpc.ClientConn
	ZKClient zkPB.ZooKeeperClient
)

func Register() bool {
	// 连接
	if !prepareConn() {
		return false
	}

	// 向zk注册
	req := &zkPB.RegisterRequest{NodeName: static.Name, NodeAddr: static.NodeAddress}
	res, err := ZKClient.Register(context.Background(), req)
	if err != nil {
		logger.Error(err)
		return false
	}

	if res.Code != "0" {
		logger.Error(res.Msg)
		return false
	}
	logger.Info(res)

	// 持续心跳
	go HeatBeating()

	return true
}

func HeatBeating() {
	// 连接
	if !prepareConn() {
		return
	}

	// 创建heartbeat request
	req := &zkPB.HeartbeatRequest{NodeName: static.Name, NodeAddr: static.NodeAddress}

	for {
		time.Sleep(time.Millisecond * 2000)
		res, err := ZKClient.Heartbeat(context.Background(), req)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		if res.Code != "0" {
			logger.Error(res.Msg)
			break
		}
		logger.Info(res)
	}
	conn.Close()
	// todo 考虑加上心跳重启机制
	logger.Info("connection to zk is closed: ", conn.GetState())
}

func UserCheckIn(userId string) bool {
	// 连接
	if !prepareConn() || strings.Trim(userId, " ") == "" {
		return false
	}

	// UserCheckIn request
	req := &zkPB.UserCheckInRequest{UserId: userId, NodeName: static.Name}

	res, err := ZKClient.UserCheckIn(context.Background(), req)
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	if res.Code != "0" {
		logger.Error(res.Msg)
		return false
	}

	return true
}

func UserCheckOut(userId string) bool {
	// 连接
	if !prepareConn() || strings.Trim(userId, " ") == "" {
		return false
	}

	// UserCheckOut request
	req := &zkPB.UserCheckOutRequest{UserId: userId, NodeName: static.Name}

	res, err := ZKClient.UserCheckOut(context.Background(), req)
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	if res.Code != "0" {
		logger.Error(res.Msg)
		return false
	}

	return true
}

func SendMessage(nodeName string, nodeAddr string, msg *nodePB.SendMessageRequest_Message) bool {
	// 连接
	newConn, err := grpc.Dial(nodeAddr, grpc.WithInsecure())
	defer newConn.Close()

	if err != nil {
		logger.Error(err)
		return false
	}

	// 创建nodeClient 去访问另外一个node 进行消息发送
	nodeClient := nodePB.NewNodeClient(newConn)

	// 创建SendMessage request
	req := &nodePB.SendMessageRequest{FromNodeName: static.Name, FromNodeAddr: static.NodeAddress, Message: msg}

	res, err := nodeClient.SendMessage(context.Background(), req)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	if res.Code != "0" {
		logger.Error(res.Msg)
		return false
	}
	logger.Info(res)
	return true
}

func FindUser(userId string) (string, string, error) {
	// 连接
	if !prepareConn() {
		return "", "", errors.New("到zookeeper的连接失败")
	}

	req := &zkPB.FindUserRequest{UserId: userId}

	res, err := ZKClient.FindUser(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}
	if res.Code != "0" {
		logger.Error(res.Msg)
		return "", "", errors.New(res.Msg)
	}

	return res.NodeName, res.NodeAddr, nil
}

func prepareConn() bool {
	if conn == nil || conn.GetState() != connectivity.Connecting {
		newConn, err := grpc.Dial(static.ZooKeeperAddress, grpc.WithInsecure())
		conn = newConn
		if err != nil {
			logger.Error(err)
			return false
		}

		// 初始化客户端
		ZKClient = zkPB.NewZooKeeperClient(conn)
	}
	return true
}
