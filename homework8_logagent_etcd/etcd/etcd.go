package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var cli *clientv3.Client

// 需要收集的日志信息的结构体
type MarshValue struct {
	Path  string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志要开往kafka中的哪个topic
}

// 初始化etcd
func Init(addr string) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second, // 超时时间
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// 从etcd中根据key获取配置项
func GetConf(key string) (M []*MarshValue, err error) { // 注意由于有多个字段，应该返回一个切片结构体
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		// 根据键key取到的值value进行反序列化   json格式的字符串 --> Go语言中的结构体变量
		err = json.Unmarshal(ev.Value, &M)
		if err != nil {
			fmt.Println("unmarshal error", err)
			return
		}
	}
	return
}

// put 设置键值
func PutConf(key string, Mvalue string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // 一秒钟计时后ctx过期，自动调用ctx.done
	_, err := cli.Put(ctx, key, Mvalue)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}

// 派哨兵监视键xxx的变化
func WatchConf(addr string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer cli.Close()
	// watch
	// 派一个哨兵一直监视着xyh这个key的值的变化 (新增、删除、修改)
	rch := cli.Watch(context.Background(), addr) // <-chan watch返回的是一个只读类型的通道

	// 从通道里尝试取监视到的信息，如果信息有改变则打印改变的情况
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// 如果键xxx的值发送改变，假如新增了一个路径，那么应该也对该路径进行读取和发送给kafka
			// 1. 新增日志
		}
	}
}

// 获取自己的本地IP地址
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("net dial error", err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
