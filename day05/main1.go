package main

import "fmt"

// 引出接口的实例
// 只要结构体实现了接口里面的方法，那么该结构体的类型也是此接口类型，就也能调用接口函数
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型
}

type cat struct{}

type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("旺旺旺")
}

// 接口函数
func f(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	f(c1)
	f(d1)
}
