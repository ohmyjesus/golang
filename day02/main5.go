package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量(长度)
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]int // 不初始化默认0值   bool类型不初始化默认false  字符串:""
	var a2 [4]int
	fmt.Printf("a1:%T  a2:%T\n", a1, a2)
	fmt.Println(a1[0])

	// 数组的初始化方式1
	a1 = [3]int{1, 2, 3}
	fmt.Println(a1[2])

	// 数组的初始化方式2	根据数组的长度自动推断数组的长度
	//a100 := [100]int{0}
	a10 := [...]int{1, 2, 3}
	fmt.Println(a10)

	// 数组的初始化方式3	根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	var a int
	fmt.Println(a)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "广州"}
	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// 2. for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	var a4 [2][3]int
	a4 = [2][3]int{
		[3]int{2, 3, 4},
		[3]int{1, 2, 3},
	}
	fmt.Println(a4)

	// 多维数组的遍历
	for _, v1 := range a4 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 3
	fmt.Println(b1[0], b2[0])

}
