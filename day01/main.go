package main

import "fmt"

// 浮点数和布尔型

func main() {
	//math.MaxFloat32 // float32位最大值
	f1 := 1.2345
	fmt.Printf("%T\n", f1) // 默认64位
	f2 := float32(1.098)
	fmt.Printf("%T\n", f2)
	// float32位的不能赋值给64位的
	//f1 = f2

	//b1 := true
	var b2 bool // 默认false
	fmt.Printf("%T value:%v\n", b2, b2)

	var s string = "hello"
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v", s)
}
