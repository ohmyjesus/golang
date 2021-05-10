package main

import (
	"fmt"
	"strings"
)

// string字符串
// 打印 F:\GoCode\src 字符串
func main() {
	//  \是有特殊含义的
	path := "'F:\\GoCode\\src'"
	fmt.Println()
	fmt.Printf("%v", path)
	fmt.Println(path)

	// 多行字符串  原样输出
	s2 := `
		123
			345
		456
	`
	fmt.Printf("%s\n", s2)

	s3 := `F:\GoCode\src`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 拼接
	name := "xyh"
	world := "ly"

	ss := name + world
	fmt.Printf("%s\n", ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s\n", ss1)

	// 分割
	ret := strings.Split(ss, "l")
	fmt.Println(ret)

	//包含  返回的是bool类型
	fmt.Println(strings.Contains(ss, "y"))

	//前缀
	fmt.Println(strings.HasPrefix(ss, "x"))

	//后缀
	fmt.Println(strings.HasSuffix(ss, "x"))

	s4 := "abcdef"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "f"))

	//将分割的字符进行拼接
	fmt.Println(strings.Join(ret, "+"))
}
