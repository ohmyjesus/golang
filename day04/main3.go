package main

import "fmt"

// 结构体
type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	// 声明一个person类型的变量
	var p person
	// 通过字段赋值
	p.name = "xyh"
	p.age = 18
	p.hobby = []string{"1", "2", "3"}
	// 访问p的字段
	fmt.Println(p.name)
	fmt.Println(p)

	// 匿名结构体 多用于一些临时场景
	var s struct {
		x string
		y int
	}
	s.x = "123"
	s.y = 18
	fmt.Println(s)
}
