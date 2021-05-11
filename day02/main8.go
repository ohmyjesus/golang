package main

import "fmt"

// make函数切片

func main() {
	a1 := make([]int, 4, 6)
	fmt.Printf("a1=%v, len(a1)=%d, cap(a1)=%d\n", a1, len(a1), cap(a1))

	a2 := make([]int, 0, 6)
	fmt.Printf("a2=%v, len(a2)=%d, cap(a2)=%d\n", a2, len(a2), cap(a2))

	// 切片的赋值
	a3 := []int{1, 3, 5}
	a4 := a3 //a3和a4都指向了同一个底层数组 切片只是一个框 真正存放的元素的是底层数组
	fmt.Println(a3, a4)
	a3[0] = 5
	fmt.Println(a3, a4)

	// 切片的遍历
	// 1. 索引遍历
	for i:=0; i < len(a3); i++{
		fmt.Println(a3[i])
	}

	// 2. for range遍历
	for i,v := range a3{
		fmt.Println(i, v)
	}
}
