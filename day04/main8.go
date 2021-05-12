package main

import "fmt"

// 给自定义类型加方法
type myInt int

func (m myInt) f() {
	fmt.Println(1)
}

func main() {
	var a myInt = 100
	a.f()
}
