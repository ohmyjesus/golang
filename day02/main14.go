package main

import "fmt"

// map 和 slice组合
func main() {
	// 普通切片
	var s1 = make([]int, 1)
	fmt.Println(s1)
	// 1. 元素类型为map的切片
	var s2 = make([]map[string]int, 3, 100)
	// 对内部的map做初始化
	s2[0] = make(map[string]int, 1)
	s2[0]["A"] = 10
	fmt.Println(s2)

	// 普通map
	var m1 = make(map[string]int, 3)
	fmt.Println(m1)
	// 2. 值为切片类型的map
	var m2 = make(map[string][]int, 3)
	m2["A"] = []int{1, 2, 3}
	fmt.Println(m2)

}
