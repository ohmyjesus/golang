package main

import "fmt"

// 闭包
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
// 把原来需要传递两个int类型的函数包装成一个不需要传参的函数
func f3(f func(int, int), x, y int) func() {
	temp := func() {
		f(x, y)
	}
	return temp
}

// 要求 f1(f2)
func main() {
	f1(f3(f2, 1, 2))
}
