package main

import (
	"fmt"
	"os"
)

// os.Args 获取命令行参数

func main() {
	fmt.Printf("%#v\n", os.Args) //06_os_args.exe a b c
	fmt.Println(os.Args[2])      // b
	fmt.Printf("%T\n", os.Args)  //[]string  string类型的切片 0 1 2 3
}
