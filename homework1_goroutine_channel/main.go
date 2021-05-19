package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 练习题 使用goroutine和channel实现一个计算int64随机数各位数和的程序。

var jobChan chan int64
var resultChan chan int64
var wg sync.WaitGroup

// 1. 开启一个`goroutine`循环生成int64类型的随机数，发送到`jobChan`
func f1(jobChan chan int64) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		randNum := rand.Int63n(100)
		jobChan <- randNum
		time.Sleep(time.Second)
	}
}

// 2. 开启24个`goroutine`从`jobChan`中取出随机数计算各位数的和，将结果发送到`resultChan`
func f2(resultChan chan int64, i int) {
	defer wg.Done()
	// 无线循环取
	for {
		res := <-jobChan
		var sum int64 = 0
		temp := res
		for res > 0 {
			remain := res % 10
			res /= 10
			sum += remain
		}
		res = temp
		resultChan <- sum
		fmt.Printf("sum计算完毕，我是%d号go线程 随机数是: %d, 求得和是: %d\n", i, res, sum)
	}
}

// 3. 主`goroutine`从`resultChan`取出结果并打印到终端输出
func main() {
	jobChan = make(chan int64, 100)
	resultChan = make(chan int64, 100)
	wg.Add(25)
	go f1(jobChan)
	for i := 1; i < 25; i++ {
		go f2(resultChan, i)
	}
	for {
		fin := <-resultChan
		fmt.Println(fin)
	}
	// Wait方法阻塞直到WaitGroup计数器减为0。
	wg.Wait()

}
