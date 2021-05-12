package main

import "fmt"

// 结构体嵌套
type person struct {
	id   int
	name string
	company
}

type company struct {
	addr string
	name string
}

func main() {
	p1 := person{
		id:   1,
		name: "xyh",
		company: company{
			addr: "cd",
			name: "tx",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.addr) // 匿名结构体，可以直接拿到最底层的字段， 但可能由于匿名结构体的相同变量，导致字段冲突
}
