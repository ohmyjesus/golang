package main

import "fmt"

// 变量的作用域
var x int = 100 //全局变量  作用域为全局

// 先在函数内部查找， 再在外部查找
// 当局部变量和全局变量冲突时， 优先使用局部变量
func f1() {
	x := 20
	fmt.Println(x)
}

func main() {
	f1()
}
