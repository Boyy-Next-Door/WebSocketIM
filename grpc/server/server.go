package grpcServer

import (
	"WebSocketIM/connection"
	pb "WebSocketIM/grpc/proto" // 引入编译生成的包
	"fmt"
	"github.com/wonderivan/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)

const (
	// todo 要通过配置文件读入
	ZooKeeperAddress = "127.0.0.1:50052" //  zookeeper的gRPC服务地址
	NodeName         = "node_01"
	NodeAddress      = "127.0.0.1：50060"
)

// 定义nodeServer并实现约定的接口
type nodeServer struct {
	pb.UnimplementedNodeServer
}

// HelloService Hello服务
var NodeServer = nodeServer{}

/**
供zookeeper调用 用于本节点某已登陆账号重复登陆 强制挤下线
*/
func (nodeServer) ForceLogout(c context.Context, r *pb.ForceLogoutRequest) (*pb.ForceLogoutResponse, error) {
	logger.Info("ForceLogout from zookeeper --- ", r.NodeName, r.UserId)

	var response *pb.ForceLogoutResponse // 校验参数
	if r.NodeName != NodeName {
		response = &pb.ForceLogoutResponse{
			Code: "1",
			Msg:  "参数节点名不正确",
		}
		return response, nil
	}

	// 本届点指定用户挤下线
	connection.Logout(r.UserId)
	response = &pb.ForceLogoutResponse{
		Code: "0",
		Msg:  "success",
	}

	return response, nil
}

func (nodeServer) SendMessage(c context.Context, r *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	logger.Info("SendMessage from "+r.FromNodeIp+":"+r.FromNodePort+" --- "+r.FromNodeName, r.Message)

	var response *pb.SendMessageResponse

	// 转换proto请求中的message为IM中的message
	message := convertMessage(r.Message)

	// 检查目标用户是否存在
	if !connection.IsOnline(message.ToUid) {
		response = &pb.SendMessageResponse{
			Code: "1",
			Msg:  "目标用户不在线，发送失败",
		}
		return response, nil
	}

	// 用户在线 向其发送消息
	err := connection.SendMessage(message.ToUid, message)

	if err != nil {
		response = &pb.SendMessageResponse{
			Code: "1",
			Msg:  err.Error(),
		}
	} else {
		response = &pb.SendMessageResponse{
			Code: "0",
			Msg:  "success",
		}
	}

	return response, nil
}

func convertMessage(msg *pb.SendMessageRequest_Message) connection.Message {
	retMsg := connection.Message{}
	retMsg.MsgId = msg.MsgId
	retMsg.MsgType = int(msg.MsgType)
	retMsg.Data = msg.Data
	retMsg.FromUid = msg.FromUid
	retMsg.ToUid = msg.ToUid
	retMsg.CreateAt = time.Unix(msg.CreateAt, 0)
	retMsg.IsRead = int(msg.IsRead)
	retMsg.ReadAt = time.Unix(msg.ReadAt, 0)
	retMsg.IsRevoke = int(msg.IsRevoke)
	retMsg.RevokeAt = time.Unix(msg.RevokeAt, 0)

	return retMsg
}

func InitGRPC() {
	listen, err := net.Listen("tcp", NodeAddress)
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterNodeServer(s, NodeServer)

	fmt.Println("gRPC server listen on " + NodeAddress)
	go s.Serve(listen)
}
