package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 功能： 实现了客户端与服务器端进行聊天的功能
// tcp的client端

func main() {
	// 与client建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}

	// 发送数据
	var msg string
	for {
		// scanln遇到空格就结束，所以采用bufio来获取输入
		fmt.Println("请说话:")
		reader := bufio.NewReader(os.Stdin)
		msg, _ = reader.ReadString('\n') // 读到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
		// 客户端读服务器端发送过来的数据
		// 定义一个缓冲区，将读到的数据放在缓冲区内，在输出到终端
		var buffer []byte = make([]byte, 128)
		n, _ := conn.Read(buffer)
		fmt.Println(string(buffer[:n]))

	}
	conn.Close()
	// 获取命令行参数，模拟实现客户端向服务器端发送不同的数据
	// if len(os.Args) < 2 {
	// 	msg = "hello world"
	// } else {
	// 	msg = os.Args[1]
	// }
}
