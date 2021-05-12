package main

import "fmt"

// fmt
func f1() {
	fmt.Print("123")
	fmt.Print("456")
	fmt.Println()

	fmt.Println("123")
	fmt.Println("456")

	// fmt.Printf("格式化字符串\n")
	// %T ：查看类型
	// %b ：二进制
	// %d ：十进制
	// %o ：八进制
	// %x ：十六进制
	// %s ：字符串
	// %c ：字符
	// %v ：值
	// %p ：指针 将指针表示为十六进制
	// %f ：浮点数
	// %t : 布尔值

	num := 90
	fmt.Printf("%d%%\n", num)

	fmt.Printf("%v\n", 65)

	// 整数 --> 字符
	fmt.Printf("%q\n", 65)
}

func f2() {
	// 获取输入
	var s1 string
	var s2 string
	fmt.Scanln(&s1, &s2)
	fmt.Println("用户输入的内容是：", s1, s2)

	var (
		name  string
		age   int
		class string
	)
	// fmt.Scanf("%s  %d  %s\n", &name, &age, &class)
	// fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}

func main() {
	var a int = 3
	ret := fmt.Sprint(a) // 将传入的数据生成为字符串并返回
	fmt.Println(ret)
	fmt.Printf("%T\n", ret)
}
