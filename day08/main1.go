package main

import (
	"fmt"
	"sync"
)

// 锁
// 用多个goroutine处理一个公共的变量 会出现问题

var x int
var wg sync.WaitGroup
var lock sync.Mutex // 定义一个锁

func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()   // 加锁
		x++           // 写操作
		lock.Unlock() // 解锁
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
