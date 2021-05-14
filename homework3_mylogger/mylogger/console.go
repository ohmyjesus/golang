package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 向对应文件内写日志相关内容
// debug 和 info 向对应文件内写日志
// warning / error / fatal 向终端中写日志

// logger 日志结构体
type Logger struct {
	flag int
}

// Newlog 日志构造函数
func (l Logger) Newlog(s string) *Logger {
	ret := stringToInt(s)
	if ret == 0 {
		panic("日志级别错误")
	}
	return &Logger{
		flag: ret,
	}
}

// 封装函数、文件和时间信息
func log(str string, st string, a ...interface{}) {
	// 生成格式化的字符串
	msg := fmt.Sprintf(st, a...)
	t := time.Now()
	funcName, fileName, line := getInfo(3) // 返回上面第3层的函数和文件信息
	// 向文件里面输出记录日志
	switch str {
	case "DEBUG":
		fileDebug, _ := os.OpenFile("F:/GoCode/src/day06/file_debug.txt", os.O_APPEND|os.O_RDWR, 0644)
		fmt.Fprintf(fileDebug, "[%s] [函数:%s] [文件名:%s] %s\t [行号:%d] %s\n", t.Format("2006/01/02 15:04:05"), funcName, fileName, str, line, msg)
	case "INFO":
		fileInfo, _ := os.OpenFile("F:/GoCode/src/day06/file_info.txt", os.O_APPEND|os.O_RDWR, 0644)
		fmt.Fprintf(fileInfo, "[%s] [函数:%s] [文件名:%s] %s\t [行号:%d] %s\n", t.Format("2006/01/02 15:04:05"), funcName, fileName, str, line, msg)
	}
	fmt.Printf("[%s] [函数:%s] [文件名:%s] %s\t [行号:%d] %s\n", t.Format("2006/01/02 15:04:05"), funcName, fileName, str, line, msg)
}

// 方法Debug 5
func (l Logger) Debug(msg string, a ...interface{}) { // ...interface{}代表可传0个或多个任意类型的参数
	// 判断是否输出
	if l.flag >= debugLevel {
		// 格式化
		log("DEBUG", msg, a...)
	} else {
		return
	}
}

// 方法Info 4
func (l Logger) Info(msg string, a ...interface{}) {
	if l.flag >= infoLevel {
		log("INFO", msg, a...)
	} else {
		return
	}
}

// 方法Warning 3
func (l Logger) Warning(msg string, a ...interface{}) {
	if l.flag >= warningLevel {
		log("WARNING", msg, a...)
	} else {
		return
	}
}

// 方法Error 2
func (l Logger) Error(msg string, a ...interface{}) {
	if l.flag >= errorLevel {
		log("ERROR", msg, a...)
	} else {
		return
	}
}

// 方法Fatal 1
func (l Logger) Fatal(msg string, a ...interface{}) {
	if l.flag >= fatalLevel {
		log("FATAL", msg, a...)
	} else {
		return
	}
}
