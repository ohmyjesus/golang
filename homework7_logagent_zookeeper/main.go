package main

import (
	"fmt"
	"logagent_zookeeper/kafka"
	taillog "logagent_zookeeper/tail_log"
	"time"

	"gopkg.in/ini.v1"
)

// 此程序为配置文件版的logagent

// 整体的思路流程为 tailf读日志 --> 将读到的日志送到通道中 -->  从通道中取值利用sarama发送到kafka中     其中读日志的配置文件用到了第三方库ini

// 注：开始运行程序之前要先启动zookeeper和kafka

var cfg *ini.File

// logagent的入口程序
func run() {
	// 1. 利用tailf读取日志
	for {
		select {
		case line := <-taillog.ReadChan(): // tails.Line为一个通道，利用通道传输消息
			// 2. 将读取到的日志发送到kafka
			kafka.SendToKafka(cfg.Section("kafka").Key("topic").String(), line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0. 加载配置文件 -- 配置文件版的logagent 用到了第三方包
	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("ini load error", err)
	}

	// 1. 初始化kafka的连接
	err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
	if err != nil {
		fmt.Println("init kafka error", err)
		return
	}
	fmt.Println("初始化kafka成功")

	// 2. 初始化tailf
	err = taillog.Init(cfg.Section("taillog").Key("path").String())
	if err != nil {
		fmt.Println("init tailf error", err)
		return
	}
	fmt.Println("初始化tailf成功")

	run()
}
