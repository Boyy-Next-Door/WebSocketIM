package fileManager

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/wonderivan/logger"
	"io"
	"os"
	"strings"
)

var FileRecordPath = "fileServer/fileManager/fileRecord.txt"
var MaxFileSizeByte = int64(500 * 1024 * 1000) // 最大支持500MB的文件
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
