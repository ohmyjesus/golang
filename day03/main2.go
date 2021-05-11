package main

import "fmt"

// defer
// defer多用于函数结束之前释放资源(文件句柄、数据库连接、socket连接)

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hhh") // 把defer后面的语句延迟执行，延迟到函数即将返回的时候再执行
	defer fmt.Println("123")
	defer fmt.Println("xyz") // 打印顺序 xyz -- 123 -- hhh
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值是x，x先被赋值为5，再对x执行++操作
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // y = x = 5 ， 再对x执行++操作，不影响y的值
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++ //改变的是形参，副本  外面的x仍然等于5
	}(x)
	return 5
}

func main() {
	deferDemo()
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
