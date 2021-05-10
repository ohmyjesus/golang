package main

import "fmt"

//整型
//占位符
//  %s %T %c %d %b %o %x 

func main() {
	//类型自动推导  十进制数
	var i1 int = 10
	fmt.Printf("%b\n", i1) //把十进制数转换为二进制
	fmt.Printf("%o\n", i1) //把十进制数转换为八进制
	fmt.Printf("%x\n", i1) //把十进制数转换为十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0xFF
	fmt.Printf("%d\n", i3)

	//查看变量的类型
	i4 := int16(9)
	fmt.Printf("%T\n", i4)

	i5 := 'c'
	fmt.Printf("%c\n", i5)
}
