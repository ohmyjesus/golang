package main

import "fmt"

// close
// 对一个已经关闭了的通道取值，取到的值是零值，取到的ok是false类型
// 如果不关闭通道，当通道已经为空时，再读就被报错deadlock!

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	<-ch1
	<-ch1
	x, ok := <-ch1 // 0 false
	fmt.Println(x, ok)
}
