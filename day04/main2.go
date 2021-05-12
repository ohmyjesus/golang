package main

import "fmt"

// 自定义类型和类型别名

// tyep后面跟的是类型
type myInt int     // 自定义类型
type yourInt = int //类型别名

func main() {
	var n myInt = 100
	fmt.Printf("%T\n", n)

	var m yourInt = 100
	fmt.Printf("%T\n", m)

	var c rune = '中'
	fmt.Printf("%T\n", c)

	// 字符串转[]byte或[]rune
	var s string = "xyh"
	ret := []byte(s)
	fmt.Println(ret)

	// []byte或[]rune转字符串
	var s1 []byte = []byte{'1', '2'}
	rret := string(s1)
	fmt.Println(rret)
}
