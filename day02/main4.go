package main

import "fmt"

// 运算符

func main() {
	// 算术运算符
	var c int = 3
	var d = 3
	var (
		a = 4
		b = 2
	)
	a++ //++ --是单独的语句，不能放在等号的右边
	b--
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(c + d)

	var e int = 5
	fmt.Println(^e)
	// 关系运算符 == != >= <= > <   返回bool类型

	//逻辑运算符 && || !			返回bool类型

	//位运算符 & | ^ << >>  与 或 异或				go语言没有 非

	//赋值运算符 = <= >= <<= >>= += -= /= %= &= |= ^=

	var aa int8 = 10 //只能存8位
	fmt.Println(aa << 10)
}
