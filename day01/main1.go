package main

import "fmt"

// 单行注释

/*
	多行注释
*/

//声明变量  推荐使用小驼峰声明法
//var name string
//var age int

//批量声明
var (
	name string
	age  int
	isok bool
)

// 关于打印
//printf 和 println 是有区别的  printf可以用后面的值替换占位符  而println是不可以用后面的值替换占位符的，会全部打印
//如 name := "xyh"  printf("%s", name) --> xyh     println("%s", name)  -->  %s xyh

func main() {
	name = "理想"
	age = 18
	isok = true
	//Go语言中变量先声明必须使用
	fmt.Printf("name:%s", name)
	fmt.Println()
	fmt.Println(isok)
	fmt.Println(age)

	//变量的初始化
	var name string = "熊宇航"
	fmt.Println(name)       //熊宇航
	fmt.Println("%s", name) //%s 熊宇航

	//类型推导
	var a = 1
	fmt.Println(a)

	//简短变量声明，只能在函数体内使用
	s3 := "ly"
	fmt.Println(s3)

	//匿名变量是一个特殊的变量
}
