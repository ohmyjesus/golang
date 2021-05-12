package main

import "fmt"

// 结构体占用一块连续的内存空间
type person struct {
	a int8 // 8个bit位 1byte
	b int8
	c int8
}

func main() {
	var p = person{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Printf("%p\n", &p.a)
	fmt.Printf("%p\n", &p.b)
	fmt.Printf("%p\n", &p.c)
}
