package main

import (
	"fmt"
	"strconv"
)

// strconv

func f1() {
	// 字符串数字 转换为 整型数字
	var s string = "123"
	i1, _ := strconv.Atoi(s)
	fmt.Printf("%#v %T\n", i1, i1) //123

	// 整型数字 转换为 字符串数字
	var i int = 100
	s1 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", s1, s1) //"100"

	// 字符串布尔值 转换为 布尔类型
	var s2 string = "true"
	b1, _ := strconv.ParseBool(s2)
	fmt.Printf("%#v %T\n", b1, b1)
}

func main() {
	// 数字97 转换为 字符串"97"
	var a int = 97
	r1 := fmt.Sprintf("%d", a)
	fmt.Printf("%#v, %T\n", r1, r1)

	// 字符串"10000" 转换为数字 10000
	var str = "10000"
	ret, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%#v, %T\n", ret, ret)

	f1()
}
