package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

// 官方库实现RPC通信

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type AreaRequest struct {
	A int
	B int
}

// 乘法运算方法
func (a *Arith) Multiply(req AreaRequest, res *int) error {
	*res = req.A * req.B
	return nil
}

func main() {
	// 1. 注册rpc服务
	rpc.Register(new(Arith))

	// 2. 采用http协议作为rpc载体
	rpc.HandleHTTP()

	fmt.Fprintf(os.Stdout, "%s\n", "start connection")

	// 3. 监听服务，等待客户端调用求面积的方法
	err := http.ListenAndServe(":8095", nil)
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
}
