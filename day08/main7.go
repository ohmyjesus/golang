package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp的客户端

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		// 向服务器端发送数据
		fmt.Print("请输入内容:")
		var data []byte = make([]byte, 1024)
		str, _ := reader.ReadString('\n')
		socket.Write([]byte(str))

		// 收到服务器端发送过来的数据并转化为大写
		n, _, _ := socket.ReadFromUDP(data)
		fmt.Println(string(data[:n]))
	}
}
