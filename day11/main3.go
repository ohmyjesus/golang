package main

import (
	"context"
	"fmt"
)

// context.WithCancel

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
				fmt.Println("我")
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) { //for range 就相当于从通道中取值
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}
