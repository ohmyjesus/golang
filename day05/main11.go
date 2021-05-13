package main

import (
	"bufio"
	"fmt"
	"os"
)

// 使用bufio获取用户输入 一直读到回车时

// 获取用户输入时如果有空格 则会出现问题
func useScan() {
	var str string
	fmt.Print("输入内容: ")
	fmt.Scanf("%s\n", &str)
	fmt.Println(str)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("输入内容: ")
	s, _ = reader.ReadString('\n')
	fmt.Println(s)

}

func main() {
	//useScan()
	useBufio()
}
