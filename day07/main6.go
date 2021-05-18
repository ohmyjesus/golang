package main

import (
	"fmt"
	"sync"
)

// channel的练习题

var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup
var once sync.Once

// 单向通道，通常用在函数的参数中
func f1(ch1 chan<- int) {
	// 1. 启动一个goroutine，生成100个数发送到ch1中
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	// 关闭通道之后不可再往通道里面发送数据，但可从通道里面发送数据出去
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	// 2. 启动一个goroutine，从ch1中取值，计算其平方放到ch2中
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		} else {
			ch2 <- x * x
		}
	}
	//close(ch2)
	once.Do(func() { close(ch2) }) // 确保某个操作只执行一次
}

func main() {
	// 3. 在main函数中，从ch2中取值并打印
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for k := range ch2 {
		fmt.Println(k)
	}

}
