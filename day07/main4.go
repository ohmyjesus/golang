package main

import (
	"fmt"
	"runtime"
	"sync"
)

// gomaxprocs
// 决定跑满CPU的几个核心数

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("a:%d\n", i)
	}

}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("b:%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)         // 默认CPU的核心数，即默认跑满CPU 如果是1表示只有一个线程来跑，会先执行a再b，或者先b在a。  如果是2,3,4... 表示多个线程执行ab，ab会同时执行
	fmt.Println(runtime.NumCPU()) // 4核8线程
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
