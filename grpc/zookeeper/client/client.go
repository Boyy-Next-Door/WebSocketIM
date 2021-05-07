package client

import (
	nodePB "WebSocketIM/grpc/node/proto"
	manager "WebSocketIM/grpc/zookeeper/nodeManager"
	"github.com/wonderivan/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func ForceLogout(nodeName string, userId string) bool {
	// 连接
	nodeAddr, err := manager.GetIns().GetNodeAddress(nodeName)
	if err != nil {
		logger.Error(err)
		return false
	}

	conn, err := grpc.Dial(nodeAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err)
		return false
	}

	// 初始化客户端
	nodeClient := nodePB.NewNodeClient(conn)

	// 强制用户下线的远程调用
	req := &nodePB.ForceLogoutRequest{NodeName: "node_01", UserId: userId}
	res, err := nodeClient.ForceLogout(context.Background(), req)
	if err != nil {
		logger.Error(err)
		return false
	}

	if res.Code != "0" {
		logger.Error(res.Msg)
		return false
	}

	return true
}
