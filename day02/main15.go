package main

import "fmt"

// 函数 既有参数也有返回值， 返回值参数可以命名，也可以不命名
// 命名的话可以直接return
func add(a int, b int) int {
	return a + b
}

// 不命名的话 return时必须加上确定变量
func add2(a int, b int) (ret int) {
	ret = a + b
	return
}

// 只有形参，没有返回值
func f1(a int, b int) {
	fmt.Println(a + b)
}

// 只有返回值，没有形参
func f2() int {
	return 2
}

// 没有返回值，也没有形参
func f3() {
	fmt.Println("111")
}

// 多个返回值
func f4() (a int, b int) {
	a = 3
	b = 4
	return
}

// 参数的类型简写
// 当形参的连续多个类型一致时，可以省略前面的类型
func f5(x, y int, m, n string) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数的最后面
func f6(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y的类型是切片
}

// Go语言中没有默认参数这个概念

func main() {
	c := add(1, 2)
	c1 := add2(1, 2)
	fmt.Println(c, c1)
	f1(2, 3)
	fmt.Println(f2())
	f3()
	_, c2 := f4()
	fmt.Println(c2)
	f6("x", 1, 2)

}
