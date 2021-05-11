package main

import "fmt"

func main() {
	// 练习题1 求数组中的和
	a5 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a5 {
		sum += v
	}
	fmt.Println(sum)

	// 练习题2 找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0, 3)和(1, 2)
	left := 0
	right := len(a5) - 1

	for left < right {
		if a5[left]+a5[right] == 8 {
			fmt.Println(left, right)
			left++
		} else if a5[left]+a5[right] < 8 {
			left++
		} else {
			right--
		}
	}
}
