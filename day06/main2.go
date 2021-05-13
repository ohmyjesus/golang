package main

import (
	"fmt"
	"time"
)

// time包

func main() {
	ret := time.Now()
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Day())
	// 时间戳
	fmt.Println(ret.Unix()) // 1620914661
	// 时间间隔 duration == time.(方法)
	fmt.Println(time.Second)
	// now + 1天
	fmt.Println(ret.Add(time.Hour * 24))
	// 定时器
	timer := time.Tick(time.Second)
	for v := range timer {
		fmt.Println(v)
	}

	// 时间格式化 2006 1 2 3 4 5
	// 把语言中的时间对象转换成字符串类型的时间
	// 2021-05-13
	fmt.Println(ret.Format("2006-01-02"))
	// 2021/05/13 14:20:20
	fmt.Println(ret.Format("2006/01/02 15:04:05"))
	// 2021/05/13 02:20:20PM
	fmt.Println(ret.Format("2006/01/02 03:04:05PM"))

	// 求两个时间的差值sub
	timerr := ret.Add(time.Hour * -1)
	tt := ret.Sub(timerr)
	fmt.Println(tt.Minutes())

	//sleep
	// 睡眠一秒
	time.Sleep(time.Second)
}
