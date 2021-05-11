package main

import "fmt"

// 函数类型
func f1() {
	fmt.Println(2)
}

func f2() int {
	return 1
}

// 函数作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	return a + b
}

// 函数作为返回类型
func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(b)
	f7 := f5(f2)
	ret := f5(f2)
	fmt.Println(ret(1, 2))
	fmt.Printf("%T\n", f7)
}
