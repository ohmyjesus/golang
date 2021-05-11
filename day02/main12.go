package main

import "fmt"

// vscode 不支持 go module

func main() {
	// 1. &取地址
	// 2. *根据地址取值
	var n int = 8
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(*p)

	// new函数申请一个内存地址
	var a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
}
