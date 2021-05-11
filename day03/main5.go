package main

import "fmt"

// 匿名函数

func main() {
	// 在函数内部是没有办法声明带名字的函数
	var f1 = func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果只是调用一次的函数，还可以写成立即调用函数
	func(x, y int) {
		fmt.Println(x + y)
	}(2, 2)

	func() {
		fmt.Println(1)
	}()

}
