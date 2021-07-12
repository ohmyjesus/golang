package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 官方库实现RPC通信

// 算数运算请求结构体
type AreaRequest struct {
	A int
	B int
}

func main() {
	// 1. 连接远程RPC服务
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := AreaRequest{9, 2}
	var res int

	// 2. 调用远程方法
	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算  Arith.Multiply 是服务器端实现的方法
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res)
}
