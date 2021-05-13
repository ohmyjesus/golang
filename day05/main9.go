package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作
func readFile1() {
	fileObj, err1 := os.Open("F:/GoCode/src/test.txt")
	if err1 != nil {
		fmt.Printf("open file failed, err %v\n", err1)
		return
	}
	fmt.Printf("%T\n", fileObj)
	// 延时关闭文件
	defer fileObj.Close()
	// 循环读文件
	for {
		var buffer = make([]byte, 2)
		n, err2 := fileObj.Read(buffer)
		if n == 0 {
			return
		}
		if err2 != nil {
			fmt.Println("read file failed")
		} else {
			fmt.Printf("读到的字节数：%d\n", n) // 字节数
			fmt.Println(string(buffer))
		}

	}
}

// 利用bufio这个包读文件
func readFile2() {
	fileObj, err1 := os.Open("F:/GoCode/src/test.txt")
	if err1 != nil {
		fmt.Printf("open file failed, err %v\n", err1)
		return
	}
	fmt.Printf("%T\n", fileObj)
	// 延时关闭文件
	defer fileObj.Close()
	// 循环读文件
	reader := bufio.NewReader((fileObj))
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 读到文件末尾
			return
		}
		if err != nil { // 读取文件失败
			fmt.Printf("read file failed, err :%v\n", err)
			return
		}
		fmt.Print(str)
	}
}

// 第三种读文件的方式
func readFile3() {
	b, err := ioutil.ReadFile("F:/GoCode/src/test.txt")
	if err != nil {
		fmt.Printf("read file failed, err %v\n", err)
	} else {
		fmt.Println(string(b))
	}
}

func main() {
	//readFile1()
	//readFile2()
	readFile3()
}
