package main

import "fmt"

// 想知道空接口中接收到的是什么类型的值

// 类型断言1
// 猜类型 一个是判断一个，一个是判断多个
func show(a interface{}) {
	str, ok := a.(string)
	if ok {
		fmt.Println("传进来的是一个字符串", str)
	} else {
		fmt.Println("猜错了")
	}
}

// 类型断言2
func show2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("this is string", t)
	case int:
		fmt.Println("this is int", t)
	case bool:
		fmt.Println("this is bool", t)
	}
}

func main() {
	var a string = "xyh"
	show(a)
	var b bool = true
	show2(b)
}
