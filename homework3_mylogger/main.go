package main

import (
	mylogger "day06/mylogger"
	"time"
)

// 测试我们自己写的日志库

func main() {
	// 声明对象
	var mylog mylogger.Logger

	for {
		// 对象调用构造函数
		log := mylog.Newlog("debug")
		id := 10010
		name := "xyh"
		log.Debug("这是一个 debug 日志 id:%d name:%s", id, name) //1 利用空接口可以传一些格式化的字符串
		log.Info("这是一个 info 日志")                           //2
		log.Warning("这是一个 warning 日志")                     //3
		log.Error("这是一个 error 日志")                         //4
		log.Fatal("这是一个 fatal 日志")                         //5
		// 睡眠三秒
		time.Sleep(time.Second * 4)
	}

}
