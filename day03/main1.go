package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1. 判断字符串中汉字的数量
func home1() {
	x1 := "hello你好"
	sum := 0
	for _, v := range x1 {
		if unicode.Is(unicode.Han, v) {
			sum++
		}
	}
	fmt.Println(sum)
}

// 2. how do you do  判断单词数量 how = 1 do = 2 you = 1
func home2() {
	x2 := "how do you do"
	// 得到一个slice
	ret := strings.Split(x2, " ")

	// 对slice进行遍历, 定义一个map
	m1 := make(map[string]int, 10)
	for _, v := range ret {
		m1[v]++
	}

	// 遍历map
	for i := range m1 {
		fmt.Println(i, m1[i])
	}
}

// 3. 回文判断
func home3() bool {
	x1 := "上海自来水来自海上"

	// 强转为rune类型
	r := []rune(x1)

	var length int = len(r) - 1
	for i := 0; i < len(r)/2; i++ {
		if r[i] == r[length-i] {
			i++
		} else {
			return false
		}
	}
	return true
}

func main() {
	home1()
	home2()
	fmt.Println(home3())
}
