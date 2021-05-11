package main

import "fmt"

// 使用copy函数复制切片

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	a3 := make([]int, 3)

	// 拷贝赋值  值拷贝
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 从切片中删除元素2
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:1], a[2:]...)
	fmt.Println(a)

	x1 := [...]int{1, 3, 5, 7}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存空间
	s1 = append(s1[:1], s1[2:]...) // 修改了底层数组！！ 前移元素覆盖掉原来的数组
	fmt.Println(s1, len(s1), cap(s1))
	// 问： x1数组元素是多少
	fmt.Println(x1, len(x1), cap(x1)) // ==> [1, 5, 7, 7]
}
