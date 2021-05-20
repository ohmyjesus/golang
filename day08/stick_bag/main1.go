package main

import (
	proto "day08/06_stick_bag/protocol"
	"fmt"
	"net"
)

// 粘包 客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		// 将消息编码解决粘包问题
		b, _ := proto.Encode(msg)
		conn.Write([]byte(b))

	}
}
