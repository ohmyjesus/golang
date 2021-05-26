package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd 的 watch 操作

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// watch
	// 派一个哨兵一直监视着xyh这个key的值的变化 (新增、删除、修改)
	rch := cli.Watch(context.Background(), "xyh") // <-chan watch返回的是一个只读类型的通道

	// 从通道里尝试取监视到的信息
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
