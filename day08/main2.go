package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
// 读共享，写独占，写的优先级高  读写锁的效率比互斥锁的效率更高

var x int = 0
var lock sync.Mutex //互斥锁
var wg sync.WaitGroup
var rwlock sync.RWMutex

// 读操作
func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.RLock() // 加读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond) // 0.001s读一次
	// lock.Unlock()
	rwlock.RUnlock() // 解读锁
}

// 写操作
func write() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock() // lock也即写锁
	x++
	time.Sleep(time.Millisecond * 5) // 0.005s写一次
	// lock.Unlock()
	rwlock.Unlock()
}

func main() {
	startTime := time.Now()
	// 写一百次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// 读一千次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime)) // 程序执行时间

}
