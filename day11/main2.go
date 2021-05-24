package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?
// 如何通知子goroutine退出

var wg sync.WaitGroup

func f2(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("world")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // Done返回的是一个只读的结构体类型的通道
			return
		default:
		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // Done返回的是一个只读的结构体类型的通道
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 以通道方式退出--官方版
	cancel() //调用此cancel函数，ctx.Done()就有了值，子goroutine就结束
	wg.Wait()
}
