package main

import (
	"fmt"
	"net"
	"strings"
)

// udp的服务端

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listenudp err", err)
		return
	}
	defer conn.Close() // 先判断listen成功，再延迟关闭

	// 不需要建立连接，直接通信
	for {
		var buffer []byte = make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("read error", err)
			return
		}
		data := buffer[:n]
		fmt.Print("收到内容")
		fmt.Println(string(data))
		// 对读到的数据转大写 返回给客户端
		str := strings.ToUpper(string(data))
		// 发送数据
		conn.WriteToUDP([]byte(str), addr)
	}
}
