package main

import "fmt"

// defer函数中的参数值都要代进去，并且要先执行内层的calc，再入栈最外层的defer函数
// 参数值被直接赋值
// 执行顺序为  calc("10", 1, 2) --> calc("20", 0, 2) --> calc("2", 0, calc("20", 0, 2)) --> calc("1", 1, calc("10", 1, 2))

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("1", x, calc("10", x, y))
	x = 0
	defer calc("2", x, calc("20", x, y))
	y = 1
}
