package main

import "fmt"

type person struct {
	name, gender string
}

// 结构体作为值类型传递
func f1(x person) {
	x.gender = "woman"
}

// 结构体作为指针类型传递
func f2(x *person) {
	//(*x).gender = "woman"
	x.gender = "woman" //语法糖
}

func main() {
	var p person
	p.name = "xyh"
	p.gender = "man"
	f1(p)
	fmt.Println(p.gender)

	f2(&p)
	fmt.Println(p.gender)

	// 结构体指针1
	var p2 = new(person)
	p2.gender = "man"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)

	// 结构体指针2
	// 2.1 key-value初始化
	var p3 = &person{
		name:   "ly",
		gender: "woman",
	}
	fmt.Println(p3)

	// 2.2 使用值列表的形式初始化
	// 值的顺序必须和结构体定义时的顺序一致
	p4 := person{
		"123",
		"45",
	}
	fmt.Println(p4)
}
