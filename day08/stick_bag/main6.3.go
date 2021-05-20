package main

import (
	"bufio"
	proto "day08/06_stick_bag/protocol"
	"fmt"
	"io"
	"net"
)

// 粘包 服务端

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// var buf [1024]byte
	for {
		// 将消息解码
		// n, err := reader.Read(buf[:])
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Println("read from client failed, err:", err)
		// 	break
		// }
		// recvStr := string(buf[:n])
		n, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		fmt.Println("收到client发来的数据：", n)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
