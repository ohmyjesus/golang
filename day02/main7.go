package main

import "fmt"

// 切片  相当于可变长度的数组

func main() {
	// 切片的定义
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // nil相当于null的意思
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"x", "y", "z"}
	fmt.Println(s1, s2)

	// 长度和容量
	fmt.Printf("len(s1): %d, cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d, cap(s2): %d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9}
	s3 := a1[0:3] // 左闭右开
	fmt.Println(s3)

	// 其他几种切片的方式
	s4 := a1[:3]  // ==> [1, 3, 5]
	s5 := a1[3:4] //  ==> [7]
	s6 := a1[:]   //  ==> [1, 3, 5, 7, 9]
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

	// 切片的容量 是指从切片的第一个元素到底层数组的最后一个元素的个数数量，只能向后拓展
	fmt.Printf("len(s4): %d, cap(s4): %d\n", len(s4), cap(s4))
	fmt.Printf("len(s5): %d, cap(s5): %d\n", len(s5), cap(s5))

	// 切片再切片  注意容量是指底层数组的元素位置
	s7 := s4[2:]
	fmt.Printf("len(s7): %d, cap(s7): %d\n", len(s7), cap(s7))

	// 切片是一个引用类型 其都指向它的底层数组
	fmt.Println(s4)
	a1[1] = 11
	fmt.Println(s4)

}
