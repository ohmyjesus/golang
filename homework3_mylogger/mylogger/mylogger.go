package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 自定义日志库
// 需求：
// 1. 可以往不同的输出位置记录日志
// 2. 日志分为5种级别    debug/ trace/ info/ warning/ error/ fatal
// 3. 日志支持开关控制   比如开发的时候什么级别的都能输出，而上线之后只有info级别以下的才能输出
// 4. 日志要有时间、行号、文件名、日志级别、日志信息

// 定义全局常量
const (
	debugLevel   = 5
	infoLevel    = 4
	warningLevel = 3
	errorLevel   = 2
	fatalLevel   = 1
)

// 将字符串全部转小写
func stringToInt(s string) int {
	str := strings.ToLower(s)
	switch str {
	case "debug":
		return 5
	case "info":
		return 4
	case "warning":
		return 3
	case "error":
		return 2
	case "fatal":
		return 1
	default:
		return 0
	}
}

// 获取行号、函数名、文件名
func getInfo(skip int) (funcName string, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("无法获取信息")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
