package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数

func f1() {
	// 方法1. flagType   传出的是一个指针
	// 创建一个标志位参数
	name := flag.String("name", "xyh", "请输入姓名") // 标志位参数name  默认值xyh  提示信息usage
	age := flag.Int("age", 20, "请输入年龄")
	married := flag.Bool("marry", true, "婚否")
	mTime := flag.Duration("time", time.Second, "结婚了多久")
	// 使用flag  07_flag.exe -name=ly -age=19 -marry=1 -time=20h
	// flag.Parse来对命令行进行解析
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*mTime) // 20h0m0s
}

func f2() {
	// 方法二. flagTypeVar  传入变量的地址进去
	var name string
	flag.StringVar(&name, "name", "xyh", "请输入姓名")
	flag.Parse()
	fmt.Println(name)
}

func main() {
	f1()
	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数 ，为[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的flag参数个数
}
