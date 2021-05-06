package main

import (
	pb "MyZooKeeper/grpc/proto" // 引入proto包
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestClient(*testing.T) {
	//start := time.Now()
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewZooKeeperClient(conn)

	req1 := &pb.RegisterRequest{NodeName: "node_02", NodeIp: "127.0.0.1", NodePort: "66666"}
	// 第一次调用
	res1, err1 := c.Register(context.Background(), req1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(res1)

	// 创建heartbeat request
	req := &pb.HeartbeatRequest{NodeName: "node_02", NodeIp: "127.0.0.1", NodePort: "66666"}

	for {
		time.Sleep(time.Millisecond * 5000)
		res, err := c.Heartbeat(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}

	//end = time.Now()
	//fmt.Println("time cost: ", end.UnixNano() - start.UnixNano())

	//// 第二次调用
	//fmt.Println(conn.GetState())
	//req.Name = "第二次调用"
	//res, err = c.SayHello(context.Background(), req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res.Message)

	// 关闭到服务器的连接
	conn.Close()
	fmt.Println(conn.GetState())

	// 尝试关闭连接后再次调用
	//fmt.Println(conn.GetState())
	//req.Name = "第三次调用"
	//res, err = c.SayHello(context.Background(), req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res.Message)

}

//func main() {
//	do()
//
//
//}
//func do() {
//	go func() {
//		for {
//			now := time.Now()
//			fmt.Println(now)
//			// 计算下一个零点
//			t := time.NewTimer(time.Second * 1)
//			<-t.C
//		}
//	}()
//}
