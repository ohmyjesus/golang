package main

import "fmt"

// switch判断
// 简化大量的判断

func main() {
	n := 2
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("de")
	}

	//在switch里定义变量，作用域不同
	switch i := 2; i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("de")
	}

	//case多个值
	switch m := 1; m {
	case 1, 4, 7:
		fmt.Println("1")
	case 2, 5, 8:
		fmt.Println("2")
	case 3, 6, 9:
		fmt.Println("3")
	default:
		fmt.Println("de")
	}

	//case后面接判断语句
	num := 2
	switch {
	case num < 1:
		fmt.Println("1")
	case num > 1 && num < 10:
		fmt.Println("2")
	case num > 10 && num < 20:
		fmt.Println("3")
	default:
		fmt.Println("de")
	}

	//fallthrough可以执行满足条件的下一个case，是为了兼容c语言中的case设计的
	num = 1
	switch num {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("de")
	}
}
