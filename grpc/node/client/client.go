package client

import (
	nodeServer "WebSocketIM/grpc/node/server"
	pb "WebSocketIM/grpc/zookeeper/proto"
	"fmt"
	"github.com/wonderivan/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
	"time"
)

var (
	conn     *grpc.ClientConn
	ZKClient pb.ZooKeeperClient
	ZKAddr   string
)

func Register() bool {
	// 连接
	conn, err := grpc.Dial(ZKAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	ZKClient := pb.NewZooKeeperClient(conn)

	// 向zk注册
	split := strings.Split(nodeServer.NodeAddress, ":")
	req1 := &pb.RegisterRequest{NodeName: "node_01", NodeIp: split[0], NodePort: split[1]}
	res1, err1 := ZKClient.Register(context.Background(), req1)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	fmt.Println(res1)
	return true
}

func HeatBeating() {
	// 创建heartbeat request
	split := strings.Split(nodeServer.NodeAddress, ":")
	req := &pb.HeartbeatRequest{NodeName: nodeServer.NodeName, NodeIp: split[0], NodePort: split[1]}

	for {
		time.Sleep(time.Millisecond * 2000)
		res, err := ZKClient.Heartbeat(context.Background(), req)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		fmt.Println(res)
	}
	conn.Close()
	// todo 考虑加上心跳重启机制
	logger.Info("connection to zk is closed: ", conn.GetState())
}
