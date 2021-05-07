package main

import (
	pb "WebSocketIM/grpc/zookeeper/proto" // 引入proto包
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
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
	c := pb.NewZooKeeperClient(conn)

	req1 := &pb.RegisterRequest{NodeName: "node_01", NodeIp: "127.0.0.1", NodePort: "12345"}
	// 第一次调用
	res1, err1 := c.Register(context.Background(), req1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(res1)

	// 创建heartbeat request
	req := &pb.HeartbeatRequest{NodeName: "node_01", NodeIp: "127.0.0.1", NodePort: "12345"}

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
//	fmt.Println(paramCheck("  3 ", "   123  "))
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

//func checkAddr(addr string) bool{
//	match, err :=  regexp.Match(`^(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\:(6553[0-5]|655[0-2]\d|65[0-4]\d{2}|6[0-4]\d{3}|[0-5]\d{4}|[1-9]\d{0,3})$`, []byte(addr))
//	if err!= nil {
//		return false
//	}
//	return match
//}
//func main() {
//	fmt.Println(checkAddr("122.168.1.1:8000"))
//}
