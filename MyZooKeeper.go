package main

//
//type GetNodeRequest struct {
//	RequestID string `json:"requestid"`
//	UserId    string `json:"userId"`
//}
//
//func chatHandler(w http.ResponseWriter, r *http.Request) {
//	t, _ := template.ParseFiles("html/im.html")
//	t.Execute(w, nil)
//}
//
//func getNodeHandler(w http.ResponseWriter, r *http.Request) {
//	logger.Info("getNode CALLED FROM --- ", r.RemoteAddr)
//
//	// 获取参数
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		logger.Error("read body err, %v\n", err)
//		return
//	}
//	logger.Info("json:", string(body))
//
//	var getNodeReq GetNodeRequest
//	if err = json.Unmarshal(body, &getNodeReq); err != nil {
//		logger.Error("Unmarshal err, %v\n", err)
//		return
//	}
//
//	// 目前支持挤下线 账号已登录不影响新的登录请求
//	// 经过负载均衡策略  将新的登录请求分配到某一个Node
//	manager := Manager.GetIns()
//	node, err := manager.GetNode(getNodeReq.UserId)
//	if err != nil {
//		ResponseUtil.InternalError(w, err.Error())
//		return
//	}
//
//	ResponseUtil.Ok(w, node, "get nodeManager success.")
//}
//
//
//func runZooKeeper() {
//	// 开启新协程 初始化grpcServer
//	zkServer.InitGRPC()
//
//	// 绑定http服务器的路由并开启服务
//	fmt.Println("server start on ", ZKhttpAddr)
//	http.HandleFunc("/chat", chatHandler)
//	http.HandleFunc("/getNode", getNodeHandler)
//	err := http.ListenAndServe(ZKhttpAddr, nil)
//	if err != nil {
//		logger.Error(err.Error())
//	}
//}
