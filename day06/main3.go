package main

import (
	"log"
	"os"
	"time"
)

// log

func main() {
	// 打开文件
	logFile, _ := os.OpenFile("F:/GoCode/src/day06/log.txt", os.O_RDWR|os.O_APPEND, 0644)
	
	// 配置日志输出位置
	log.SetOutput(logFile)
	for {
		log.Println("this is test log")
		time.Sleep(time.Second * 3)
	}
}
