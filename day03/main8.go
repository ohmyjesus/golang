package main

import "fmt"

// panic和recover
// recover()必须搭配defer使用
func f1() {
	fmt.Println("1")
}

func f2() {
	// 刚刚连接数据库
	defer func() {
		err := recover() // 尝试修复
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现错误") // 程序崩溃退出
	fmt.Println("2")
}

func f3() {
	fmt.Println("3")
}

func main() {
	f1()
	f2()
	f3()
}
