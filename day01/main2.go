package main

import "fmt"

//常量
const pi = 3.1415

//批量声明常量
const (
	code1 = 200
	code2 = 403
	code3 = 404
)

// 如果某一行没有写值，则后面的默认和前面的一样
const (
	n1 = 100
	n2
	n3
)

// iota 类似枚举  在const关键字出现时将被重置为0,每新增一行常量声明则iota值加1
const (
	a1 = iota // 0
	a2 = iota // 1
	_         // 2
	a3        // 3
)

// 插队
const (
	b1 = iota //0
	b2 = 100  //1
	b3 = iota //2
	b4        //3
)

// 多个变量声明在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main() {
	//fmt.Println(n2)
	fmt.Println(a3)
	fmt.Println(b4)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
