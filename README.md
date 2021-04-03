#	需要重新整理单机版的后台逻辑，把分包规范化一下，现在感觉太分散了。
#	写一个简单的zookeeper（http-server && grpc-server），功能如下：
## grpc
- func register(ip,port,nodeId) --- 每一个node开启后需要向zk注册，提交自己的ip、port，即调用zookeeper的register方法
- func heartbeat(ip,port,nodeId) --- 每一个node在注册之后，都需要定时向zk发送心跳，当zk检测到某个node失活，则会丢入一个淘汰队列等待其复活，未复活成功则清空当前map中原本存在于该node上的所有用户
- func addUser(userId,nodeId) --- 当zk给一个sdk分配了一个node时，zk不会记录用户的登录位置，当sdk真正连接上node时，再由node主动告知zk该用户登录
- func dropUser(userId,nodeId) --- 当与node创建websocket连接的sdk断开连接，那么node主动告知zk
- func findUser(userId) --- 当node的ws连接报文中涉及到一个不存在本node的userId,则向zk询问对方位置，如果没有返回有效nodeId，认为该用户不在线，反之该node调用目标node的send方法
## http 
- POST GetLoginNode --- SDK要首先通过http请求向zk请求一个node节点
- GET /chat	--- 返回聊天前端页面
- GET /admin  --- 返回后台管理前端页面 （可以简单查看当前zk信息、node列表、node负载、调整负载均衡策略、开关node、查看聊天记录等）
#	实现node节点（grpc-server && websocket-server），功能如下：
## websocket
- ws://ip:port/ws 建立websocket连接
-	websocket内部报文：	
```golang
  	SEND       = 1
  	REVOKE     = 2
	LOGIN      = 3
	LOGOUT     = 4
	READ_ACK   = 5
	SENT       = 6
	REVOKE_ACK = 7
```
## grpc
-	func send() --- 跨node时，由一方调用另一方的send()方法，实现ws报文的传递。
