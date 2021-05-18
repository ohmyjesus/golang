package main

import (
	"fmt"
	"time"
)

// goroutine线程

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		// go hello(i) // 开启一个单独的goroutine去执行hello函数(任务)
		// 匿名函数
		go func(i int) {
			fmt.Printf("%d --- \n", i)
		}(i)
	}
	fmt.Println("main")
	// main结束了, 由main函数启动的goroutine也就结束了
	// 等待1秒
	time.Sleep(time.Second)
}
