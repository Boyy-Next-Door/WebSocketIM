package main

import (
	pb "WebSocketIM/grpc-test/proto" // 引入proto包
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	//start := time.Now()
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)
	// 创建request
	req := &pb.HelloRequest{Name: "gRPC"}
	// 第一次调用
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
	//end = time.Now()
	//fmt.Println("time cost: ", end.UnixNano() - start.UnixNano())

	// 第二次调用
	fmt.Println(conn.GetState())
	req.Name = "第二次调用"
	res, err = c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
	// 关闭到服务器的连接
	conn.Close()
	fmt.Println(conn.GetState())

	// 尝试关闭连接后再次调用
	fmt.Println(conn.GetState())
	req.Name = "第三次调用"
	res, err = c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)

}
