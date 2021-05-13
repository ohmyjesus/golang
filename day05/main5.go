package main

import "fmt"

// 空接口  没有必要写空接口的名字
// 任何类型的结构体都可以调用接口函数
// 空接口可以存储任意类型的数据

type dog struct {
	name string
}

// 空接口作为函数参数
func show(i interface{}) {
	fmt.Println("this is empty interface")
}

func main() {
	m1 := make(map[string]interface{})
	m1["name"] = "xyh"
	m1["age"] = 17
	m1["gender"] = true
	m1["hobby"] = [...]string{"1", "2", "3"}
	show(m1) // 传入map类型到空接口
	fmt.Println(m1)

	var d1 = dog{"tm"}
	show(d1) // 传入dog结构体类型到空接口
}
