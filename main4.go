package main

// runtime.Caller()
// 获取行号、文件名、函数名

import (
	"fmt"
	"path"
	"runtime"
)

func getInfo(skip int) {
	pc, file, line, ok := runtime.Caller(skip) // skip为0 表示返回调用当前函数的函数名和文件名
	if !ok {
		fmt.Println("无法获取信息")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)            // 04_runtime/main.go 函数名
	fmt.Println(path.Base(file)) // 获取最后一个路径 main.go
	fmt.Println(line)            // 行号
}

func main() {
	var skip int = 0
	getInfo(skip) // skip为1表示上一层调用此封装函数的函数名和文件名
}
