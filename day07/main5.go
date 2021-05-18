package main

import (
	"fmt"
	"sync"
)

// channel
// 分为带缓冲区的channel和不带缓冲区的channel

var wg sync.WaitGroup

var c chan int // 需要指定通道中元素的类型

func noBufChannel() {
	c = make(chan int) //无缓冲区的通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-c // 取出数据
		fmt.Println("go routine", x)
	}()

	c <- 10 // 发送到通道中  缓冲区为0的不能接收数据，需要启动一个线程来取出数据，否则会阻塞等待
	fmt.Println("10发送到了通道c中")
	// Wait方法阻塞直到WaitGroup计数器减为0。
	wg.Wait()
	close(c)
}

func bufChannel() {
	c = make(chan int, 10) // 缓冲区为10的通道的初始化
	c <- 10                // 发送到通道中
	fmt.Println("10发送到了通道c中")
	x := <-c // 从一个通道中接收值
	fmt.Println(x)
}

func main() {
	noBufChannel()
	//bufChannel()
}
