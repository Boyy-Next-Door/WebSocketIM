package main

import (
	"WebSocketIM/connection"
	"WebSocketIM/consumer"
	"WebSocketIM/datasource"
	"WebSocketIM/mq"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

// 处理/upload 逻辑
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

func chatHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./im.html")
	t.Execute(w, nil)
}
func historyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	method := r.Method
	log.Println(method, ": getHistory ", r.RemoteAddr) //获取请求的方法

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
func main() {
	datasource.InitDB()
	mq.Init()
	consumer.Init()
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/chat", chatHandler)
	http.HandleFunc("/getHistory", historyHandler)
	err := http.ListenAndServe("0.0.0.0:7777", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	//mongodb.Init()
}
