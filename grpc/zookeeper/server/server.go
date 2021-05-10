package grpcServer

import (
	Manager "WebSocketIM/grpc/zookeeper/nodeManager"
	pb "WebSocketIM/grpc/zookeeper/proto" // 引入编译生成的包
	"WebSocketIM/static"
	"fmt"
	"github.com/wonderivan/logger"
	_ "github.com/wonderivan/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"strings"
)

// 定义helloService并实现约定的接口
type zkServer struct {
	pb.UnimplementedZooKeeperServer
}

/**
供node远程调用的注册接口
*/
func (zkServer) Register(c context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	logger.Info("register from ", r.HttpAddr, " --- ", r.NodeName)

	manager := Manager.GetIns()
	newNode := Manager.Node{NodeName: r.NodeName, GrpcAddr: r.GrpcAddr, HttpAddr: r.HttpAddr}
	err := manager.AddNode(newNode)

	var response *pb.RegisterResponse
	if err != nil {
		response = &pb.RegisterResponse{
			Code: "1",
			Msg:  "register failed --- " + err.Error(),
		}
	} else {
		response = &pb.RegisterResponse{
			Code: "0",
			Msg:  "register success.",
		}
	}
	return response, nil
}

/**
供node调用的心跳接口
*/
func (zkServer) Heartbeat(c context.Context, r *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	logger.Info("heartbeat from ", r.HttpAddr, " --- ", r.NodeName)

	manager := Manager.GetIns()
	err := manager.HeartBeat(r.NodeName)

	var response *pb.HeartbeatResponse
	if err != nil {
		response = &pb.HeartbeatResponse{
			Code: "1",
			Msg:  "heart beat failed --- " + err.Error(),
		}
	} else {
		response = &pb.HeartbeatResponse{
			Code: "0",
			Msg:  "heart beat success.",
		}
	}
	return response, nil
}

/**
供node调用的用户登入接口
*/
func (zkServer) UserCheckIn(c context.Context, r *pb.UserCheckInRequest) (*pb.UserCheckInResponse, error) {
	logger.Info("UserCheckIn from ", r.NodeName, " --- ", r.UserId)
	var response *pb.UserCheckInResponse

	// 参数校验
	if isEmptyStr(r.UserId) || isEmptyStr(r.NodeName) {
		response = &pb.UserCheckInResponse{
			Code: "1",
			Msg:  "user check in failed --- " + "参数错误",
		}
		return response, nil
	}

	manager := Manager.GetIns()
	err := manager.UserCheckIn(r.NodeName, r.UserId)

	if err != nil {
		response = &pb.UserCheckInResponse{
			Code: "1",
			Msg:  "user check in failed --- " + err.Error(),
		}
	} else {
		response = &pb.UserCheckInResponse{
			Code: "0",
			Msg:  "user check in success.",
		}
	}
	return response, nil
}

/**
供node调用的用户登出接口
*/
func (zkServer) UserCheckOut(c context.Context, r *pb.UserCheckOutRequest) (*pb.UserCheckOutResponse, error) {
	logger.Info("UserCheckOut from ", r.NodeName, " --- ", r.UserId)
	var response *pb.UserCheckOutResponse

	// 参数校验
	if isEmptyStr(r.UserId) || isEmptyStr(r.NodeName) {
		response = &pb.UserCheckOutResponse{
			Code: "1",
			Msg:  "user check out failed --- " + "参数错误",
		}
		return response, nil
	}

	manager := Manager.GetIns()
	err := manager.UserCheckOut(r.NodeName, r.UserId)

	if err != nil {
		response = &pb.UserCheckOutResponse{
			Code: "1",
			Msg:  "user check out failed --- " + err.Error(),
		}
	} else {
		response = &pb.UserCheckOutResponse{
			Code: "0",
			Msg:  "user check out success.",
		}
	}

	return response, nil
}

/**
供node调用的用户查询接口
*/
func (zkServer) FindUser(c context.Context, r *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	logger.Info("FindUser from ", " --- ", r.UserId)
	var response *pb.FindUserResponse

	// 参数校验
	if isEmptyStr(r.UserId) {
		response = &pb.FindUserResponse{
			Code: "1",
			Msg:  "find user failed --- " + "参数错误",
		}
		return response, nil
	}

	manager := Manager.GetIns()
	node, err := manager.FindUser(r.UserId)

	if err != nil {
		response = &pb.FindUserResponse{
			Code: "1",
			Msg:  "find user failed --- " + err.Error(),
		}
	} else {
		response = &pb.FindUserResponse{
			Code:     "0",
			Msg:      "find user success.",
			NodeName: node.NodeName,
			GrpcAddr: node.GrpcAddr,
			HttpAddr: node.HttpAddr,
		}
	}

	return response, nil
}

func isEmptyStr(str string) bool {
	return strings.Trim(str, " ") == ""
}

// HelloService Hello服务
var ZooKeeperServer = zkServer{}

func InitGRPC() {
	listen, err := net.Listen("tcp", static.ZooKeeperAddress)
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterZooKeeperServer(s, ZooKeeperServer)

	fmt.Println("gRPC server listen on " + static.ZooKeeperAddress)
	go s.Serve(listen)
}

//func main() {
//	listen, err := net.Listen("tcp", ZKgrpcAddr)
//	if err != nil {
//		fmt.Println("Failed to listen: %v", err)
//	}
//
//	// 实例化grpc Server
//	s := grpc.NewServer()
//
//	// 注册HelloService
//	pb.RegisterZooKeeperServer(s, ZooKeeperServer)
//
//	fmt.Println("Listen on " + ZKgrpcAddr)
//	s.Serve(listen)
//}
