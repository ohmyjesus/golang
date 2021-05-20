package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 功能： 实现了客户端与服务器端进行聊天的功能
// tcp的server端

func process(conn net.Conn, i int) {
	defer conn.Close()
	// 3. 与客户端通信
	// 线程循环处理
	for {
		var temp []byte = make([]byte, 128)
		n, err := conn.Read(temp) // 没有数据传来则阻塞等
		if err != nil {
			fmt.Println("read error", err)
			return
		}
		// 将客户端发送过来的数据输出到终端
		fmt.Println(string(temp[:n]))

		// 服务器端回复客户端
		fmt.Println("请回复:")
		var answer string
		reader := bufio.NewReader(os.Stdin)
		answer, _ = reader.ReadString('\n')
		conn.Write([]byte(answer))
		fmt.Println()
	}
}

func main() {
	// 1. 本地端口启动服务
	var i int = 1
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen error", nil)
		return
	}
	defer listener.Close()
	for {
		// 2. 等待客户端连接，没有连接则阻塞等待，一旦连接上，则启动一个线程来做数据交互
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accepy error", err)
			return
		}
		go process(conn, i)
		i++
	}
}
