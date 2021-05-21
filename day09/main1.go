package main

import (
	"fmt"
	"io"
	"net/http"
)

// http client

func main() {
	resp, err := http.Get("http://192.168.123.176:8080/ly?name=123&age=456") // 模拟浏览器发送http请求
	if err != nil {
		fmt.Println("get url error ", err)
		return
	}
	defer resp.Body.Close() // 记得关闭body
	// 从resp中把服务器返回的数据读出来

	// var data []byte
	// resp.Body.Read(data)
	// resp.Body.Close()

	b, err := io.ReadAll(resp.Body) // 在客户端读出服务器端返回响应的body
	if err != nil {
		fmt.Println("read body error ", err)
		return
	}
	fmt.Println(string(b))
}
