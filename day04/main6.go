package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数: 约定俗成 用new开头
// 返回的是结构体还是结构体指针
// 当结构体大时，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("xyh", 18)
	p2 := newPerson("ly", 20)
	fmt.Println(p1, p2)
}
