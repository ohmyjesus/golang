package main

// 包的路径
// 想要被的别的包调用的标识符都要首字母大写
// 单行导入、多行导入
// 导入包的时候可以指定别名
// 可以使用匿名导入包
// 每个包导入的时候会自动执行一个名为init()函数，它没有参数也没有返回值，也不能手动调用
// 多个包中都定义了init()函数、则执行顺序为压栈出栈顺序

// 执行顺序： 全局声明 --> init函数 --> main函数

import (
	calc "day05/07_package"
	"fmt"
)

var a int = 3

func init() {
	fmt.Println("this is package main8.go init")
	fmt.Println(a)
}

func main() {
	ret := calc.Add(2, 3)
	fmt.Println(ret)
}
