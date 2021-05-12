package main

import "fmt"

// 匿名字段
// 适用于场景适用于字段比较简单
// 不常用

type person struct {
	int
	string
	//隐含 string : string
}

func main() {
	p1 := person{
		10,
		"n",
	}
	fmt.Println(p1.int)
	fmt.Println(p1.string)
}
