package main

import "fmt"

// 递归
// 阶乘
func f(n int) int {
	if n > 1 {
		ret := n * f(n-1)
		return ret
	} else {
		return 1
	}
}

func main() {
	ret := f(5)
	fmt.Println(ret)
}
