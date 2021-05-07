package main

import (
	"WebSocketIM/connection"
	"WebSocketIM/datasource"
	nodeClient "WebSocketIM/grpc/node/client"
	nodeServer "WebSocketIM/grpc/node/server"
	Manager "WebSocketIM/grpc/zookeeper/nodeManager"
	zkServer "WebSocketIM/grpc/zookeeper/server"
	"WebSocketIM/mq"
	"WebSocketIM/util"
	ResponseUtil "WebSocketIM/util"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wonderivan/logger"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// todo 从配置文件中读
// 所有节点（包括zk和node两种类型） 都需要占用两个端口
// 对于zookeeper: 一个端口用于开启http服务 提供聊天页面访问、获取node、服务监控等功能 / 一个用于开启gRPC server 供node远程调用
// 对于node：	 一个端口用于开启http服务 供sdk升级成websocket连接、通过http请求拉取聊天记录  / 一个用于开启gRPC server 供zookeeper远程调用
var (
	NodeHttpAddress = ""
	ZKhttpAddr      = ""
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		conn   *connection.Connection
		//data []byte
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	//创建连接，开启携程监听ws中的收发消息，
	if conn, err = connection.InitConnection(wsConn); err != nil {
		goto ERR
	}

	return
ERR:
	//conn.WriteMessage(([]byte)("error:"+err.Error()))
	conn.Close()

}

type GetNodeRequest struct {
	RequestID string `json:"requestid"`
	UserId    string `json:"userId"`
}

// 上传文件接口
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在upload目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

// 获取历史消息接口
func historyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	method := r.Method
	logger.Info(method, ": getHistory ", r.RemoteAddr) //获取请求的方法

	data := r.URL.Query()
	userId := data.Get("userId")
	fmt.Println(userId)
	targetId := data.Get("targetId")
	fmt.Println(targetId)
	timeUnix, err := strconv.Atoi(data.Get("timeBefore"))
	if err != nil {
		w.Write([]byte("parameter error."))
		return
	}
	timeBefore := time.Unix(int64(timeUnix), 0)
	fmt.Println(timeBefore)

	num, err := strconv.Atoi(data.Get("num"))
	if err != nil {
		w.Write([]byte("parameter error."))
		return
	}
	fmt.Println(num)

	ret := make([]connection.Message, 0)

	//装载数据
	datasource.Select(datasource.SelectHistory, &ret, userId, targetId, timeBefore, num, targetId, userId, timeBefore, num, num)

	marshal, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}

// 访问聊天页面接口
func chatHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/im.html")
	t.Execute(w, nil)
}

// 请求node接口
func getNodeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("getNode CALLED FROM --- ", r.RemoteAddr)

	// 获取参数
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("read body err, %v\n", err)
		return
	}
	logger.Info("json:", string(body))

	var getNodeReq GetNodeRequest
	if err = json.Unmarshal(body, &getNodeReq); err != nil {
		logger.Error("Unmarshal err, %v\n", err)
		return
	}

	// 目前支持挤下线 账号已登录不影响新的登录请求
	// 经过负载均衡策略  将新的登录请求分配到某一个Node
	manager := Manager.GetIns()
	node, err := manager.GetNode(getNodeReq.UserId)
	if err != nil {
		ResponseUtil.InternalError(w, err.Error())
		return
	}

	ResponseUtil.Ok(w, node, "get nodeManager success.")
}

func main() {
	// 读取命令行参数
	mode := ""
	name := ""
	httpAddr := ""
	grpcAddr := ""
	zkAddr := ""
	flag.StringVar(&mode, "m", "node", "运行模式： node / zookeeper 默认为前者")
	flag.StringVar(&name, "n", "undefined", "节点名：当mode为node时，这是到zk注册的唯一标识，不能重复")
	flag.StringVar(&httpAddr, "h", "", "http服务开启的ip和端口号")
	flag.StringVar(&grpcAddr, "g", "", "gRPC服务开启的ip和端口号")
	flag.StringVar(&zkAddr, "z", "", "注册中心地址：当mode为node时 必须传入此参数")
	flag.Parse()

	// 校验参数
	if mode != "node" && mode != "zookeeper" || name == "undefined" || !util.CheckAddr(httpAddr) || !util.CheckAddr(grpcAddr) || mode == "node" && !util.CheckAddr(zkAddr) {
		logger.Error("参数校验失败")
		return
	}

	//根据启动的模式装载server参数
	if mode == "node" {
		// node模式
		NodeHttpAddress = httpAddr
		nodeServer.NodeAddress = grpcAddr
		nodeServer.NodeName = name
		nodeClient.ZKAddr = zkAddr
		// 启动服务
		runNode()
	} else {
		// zookeeper模式
		ZKhttpAddr = httpAddr
		zkServer.ZKgrpcAddr = grpcAddr
		zkServer.ZKname = name

		// 启动服务
		runZooKeeper()
	}
}

func runNode() {
	// 初始化db
	datasource.InitDB()
	// 初始化消息队列
	mq.Init()
	// 初始化消息队列的消费者
	connection.InitConsumer()
	// 初始化gRPC服务端
	nodeServer.InitGRPC()
	// 注册
	success := nodeClient.Register()
	if !success {
		logger.Error("注册失败")
		return
	}
	// 持续心跳
	go nodeClient.HeatBeating()

	// 绑定http服务器路由并开启http服
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/getHistory", historyHandler)
	err := http.ListenAndServe(NodeHttpAddress, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	//mongodb.InitConsumer()
}

func runZooKeeper() {
	// 开启新协程 初始化grpcServer
	zkServer.InitGRPC()

	// 绑定http服务器的路由并开启服务
	fmt.Println("server start on ", ZKhttpAddr)
	http.HandleFunc("/chat", chatHandler)
	http.HandleFunc("/getNode", getNodeHandler)
	err := http.ListenAndServe(ZKhttpAddr, nil)
	if err != nil {
		logger.Error(err.Error())
	}
}
