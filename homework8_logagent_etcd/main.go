package main

import (
	"fmt"
	"logagent_etcd/etcd"
	"logagent_etcd/kafka"
	taillog "logagent_etcd/tail_log"
	"sync"
	"time"

	"github.com/hpcloud/tail"
	"gopkg.in/ini.v1"
)

// 此程序为etcd版的logagent

// 整体的思路为 初始化kafka和etcd，连接上etcd之后不再从配置文件里读，而是从etcd里面获取日志信息，然后创建子goroutine对日志信息进行监视并将日志发送到kafka，遗憾的是未能实现日志的动态新增功能。

// 注：开始运行程序之前要先启动etcd.exe和kafka 同时启动两个kafka的消费者将日志信息打印到终端上，此时变更log日志信息时，可以在终端上看到日志信息的变化

// logagent的入口程序
func run(cfg *ini.File, tails *tail.Tail, v *etcd.MarshValue) {
	defer wg.Done()
	// 1. 利用tailf读取日志
	for {
		select {
		case line := <-taillog.ReadChan(tails): // tails.Line为一个通道，利用通道传输消息
			// 2. 将读取到的日志发送到kafka 和 打印到终端
			fmt.Println(line.Text)
			kafka.SendToKafka(v.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

var wg sync.WaitGroup

func main() {
	// 0. 加载配置文件 -- 用到了第三方包
	var cfg *ini.File
	cfg, _ = ini.Load("./conf/config.ini")

	// 1. 初始化kafka的连接
	err := kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
	if err != nil {
		fmt.Println("init kafka error", err)
		return
	}
	fmt.Println("初始化kafka成功")

	// 2. etcd取代tailf
	// 初始化etcd
	err = etcd.Init(cfg.Section("etcd").Key("address").String())
	if err != nil {
		fmt.Println("etcd init error", err)
		return
	}
	fmt.Println("初始化etcd成功")

	// 3. 设置键值对传给etcd  值为json格式的字符串
	str := `[{"path":"c:/tmp/nginx.log","topic":"nginx_log"},{"path":"d:/xxx/redis.log","topic":"redis_log"}]`

	// 为了实现每个logagent都拉取到自己独有的配置，所以要以自己的IP地址作为区分
	ip, _ := etcd.GetOutboundIP()
	addr := fmt.Sprintf(cfg.Section("etcd").Key("collect_log_key").String(), ip) // sprintf加了IP的地址

	// 添加的键值对为 key:   /logagent/192.168.123.176/collect_config   value:    [{"path":"c:/tmp/nginx.log","topic":"nginx_log"},{"path":"d:/xxx/redis.log","topic":"redis_log"}]
	etcd.PutConf(addr, str)

	// 4. 连接上etcd之后不再从配置文件里读，而是从etcd里面获取日志信息
	M, err := etcd.GetConf(addr)
	if err != nil {
		fmt.Println("get conf error", err)
		return
	}

	// 获得了日志收集的路径和日志应该发往的topic  c:/tmp/nginx.log-->nginx__log  2. d:/xxx/redis.log-->redis_log
	// 5. 对每一个路径都启动goroutine去调用watch和run函数
	for _, v := range M {
		wg.Add(2)
		// 5.1 派哨兵去监视日志收集项的变化
		go etcd.WatchConf(addr, &wg)

		// 5.2 去路径下读取日志并发往kafka
		tails, err := taillog.Init(v.Path)
		if err != nil {
			fmt.Println("init tailf error", err)
			return
		}
		fmt.Println("初始化tailf成功")

		// 5.3 每一个日志项都起一个gorouinte执行
		go run(cfg, tails, v)
	}
	wg.Wait()

}
