package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?
// 如何通知子goroutine退出

var wg sync.WaitGroup
var flag bool
var c chan int

func f() {
	defer wg.Done()
	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
		if flag {
			break
		}
		// 利用select操作取通道里的值
		select {
		case <-c:
			return // 返回之前执行defer
		default:
		}
	}
}

// 1. 全局变量的方式通知退出
func solve1() {
	flag = true
}

// 2. 以通道的形式退出
func solve2() {
	defer wg.Done()
	n := 3
	c <- n
}

func main() {
	wg.Add(2)
	c = make(chan int, 1) // 缓冲区的大小为1
	go f()
	time.Sleep(time.Second * 5)
	// go solve1()
	go solve2()
	wg.Wait()
	close(c)
}
