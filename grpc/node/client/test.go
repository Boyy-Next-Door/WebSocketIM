package client

func main() {
	////start := time.Now()
	//// 连接
	//conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer conn.Close()
	//
	//// 初始化客户端
	//c := pb.NewZooKeeperClient(conn)
	//
	// 查询用户
	//req1 := &pb.FindUserRequest{NodeName: "node_01", UserId: "杨噶劲"}
	//res1, err1 := c.FindUser(context.Background(), req1)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//fmt.Println(res1)
	//
	//// 用户登陆
	//req2 := &pb.UserCheckInRequest{NodeName: "node_01", UserId: "杨噶劲"}
	//res2, err2 := c.UserCheckIn(context.Background(), req2)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//fmt.Println(res2)
	//
	//// 再次查询
	//res1, err1 = c.FindUser(context.Background(), req1)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//fmt.Println(res1)
	//
	//// 用户登出
	//req3 := &pb.UserCheckOutRequest{NodeName: "node_01", UserId: "杨噶劲"}
	//res3, err3 := c.UserCheckOut(context.Background(), req3)
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//fmt.Println(res3)
	//
	//// 再次查询
	//res1, err1 = c.FindUser(context.Background(), req1)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//fmt.Println(res1)
}
