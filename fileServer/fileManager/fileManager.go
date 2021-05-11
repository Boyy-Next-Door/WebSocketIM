package fileManager

import (
	zkPB "WebSocketIM/grpc/zookeeper/proto"
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/wonderivan/logger"
	"google.golang.org/grpc"
	"io"
	"os"
	"strings"
)

var (
	conn            *grpc.ClientConn
	ZKClient        zkPB.ZooKeeperClient
	FileRecordPath  = "fileServer/fileManager/fileRecord.txt"
	MaxFileSizeByte = int64(500 * 1024 * 1000) // 最大支持500MB的文件
)

// 缓存当前fileMapping中记录的文件映射关系
var fileIdSet = mapset.NewSet() // 文件id集合

func Init() {
	// 初始化文件映射中心 从fileMapping中读取当前存储的文件
	filepath := FileRecordPath
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				logger.Info("fileRecord.txt load success!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		line = strings.TrimSpace(line)
		// 改行为注释行
		if line[0] == '#' {
			continue
		} else {
			fileIdSet.Add(line)
		}
	}
}

func SaveFile(fileId string) error {
	fileIdSet.Add(fileId)
	filepath := FileRecordPath
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return err
	}
	defer file.Close()

	// 查找文件末尾的偏移量
	n, _ := file.Seek(0, 2)
	// 从末尾的偏移量开始写入内容
	_, err = file.WriteAt([]byte(fileId+"\n"), n)
	return err
}

//func Register() bool {
//	// 连接
//	if !prepareConn() {
//		return false
//	}
//
//	// 向zk注册
//	req := &zkPB.RegisterRequest{NodeName: static.Name, HttpAddr: static.HttpAddress, GrpcAddr: static.GrpcAddress, NodeType: static.Mode}
//	res, err := ZKClient.Register(context.Background(), req)
//	if err != nil {
//		logger.Error(err)
//		return false
//	}
//
//	if res.Code != "0" {
//		logger.Error(res.Msg)
//		return false
//	}
//	logger.Info(res)
//
//	// 持续心跳
//	go HeatBeating()
//
//	return true
//}
//
//func HeatBeating() {
//	// 连接
//	if !prepareConn() {
//		return
//	}
//
//	// 创建heartbeat request
//	req := &zkPB.HeartbeatRequest{NodeName: static.Name, GrpcAddr: static.GrpcAddress, HttpAddr: static.HttpAddress, NodeType: static.Mode}
//
//	for {
//		time.Sleep(time.Millisecond * 2000)
//		res, err := ZKClient.Heartbeat(context.Background(), req)
//		if err != nil {
//			logger.Error(err.Error())
//			break
//		}
//		if res.Code != "0" {
//			logger.Error(res.Msg)
//			break
//		}
//		//logger.Info(res)
//	}
//	conn.Close()
//	// todo 考虑加上心跳重启机制
//	logger.Info("connection to zk is closed: ", conn.GetState())
//}
//func prepareConn() bool {
//	if conn == nil || conn.GetState() != connectivity.Connecting {
//		newConn, err := grpc.Dial(static.ZooKeeperAddress, grpc.WithInsecure())
//		conn = newConn
//		if err != nil {
//			logger.Error(err)
//			return false
//		}
//
//		// 初始化客户端
//		ZKClient = zkPB.NewZooKeeperClient(conn)
//	}
//	return true
//}
