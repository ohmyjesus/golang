package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup 等待所有goroutine执行完之后再退出

func f() {
	rand.Seed(time.Now().UnixNano()) // 加随机数种子，保证每次执行的时候都不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(11) // [0,11)
		fmt.Println(r1, r2)
	}

}

func f1(i int) {
	// 睡眠一个随机数
	defer wg.Done() // 执行完goroutine之后计数器减1
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

// 等待所有的goroutine执行完之后再退出
var wg sync.WaitGroup

func main() {
	f()
	//wg.Add(10)  表示启用几个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1) //执行goroutine之前计数器加1
		go f1(i)
	}
	// 如何知道这十个goroutine都结束了
	wg.Wait() // 计数器减为0时则程序结束
}
