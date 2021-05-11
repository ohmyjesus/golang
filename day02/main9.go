package main

import "fmt"

// append() 为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "广州"}
	fmt.Println(s1)
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素时，原来的底层数组放不下的时候，Go底层就会换一个新的底层数组，为了防止原来的切片丢失，需要用原来的切片来进行接收
	s1 = append(s1, "成都")
	fmt.Println(s1)
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))

	// append追加多个元素
	s2 := []string{"大连", "重庆", "西安"}
	s1 = append(s1, s2...) //...表示拆开
	fmt.Println(s1)
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))

}
