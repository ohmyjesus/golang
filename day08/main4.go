package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	// x++ // x = x + 1
	atomic.AddInt64(&x, 1) // 利用原子操作对值加1  在内部也实现了加锁和解锁   可以直接调用这个即可，方便调用
}

func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
