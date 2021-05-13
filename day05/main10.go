package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写文件
func writeFile1() {
	// O_APPEND 追加  O_TRUNC 每次写都清空
	fileObj, err1 := os.OpenFile("./src/day05/10_file_write/test.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Printf("open file error :%v\n", err1)
		return
	}
	var myByte []byte = []byte("come on\n")
	// write
	n, err2 := fileObj.Write(myByte)
	if err2 != nil {
		fmt.Printf("write file error :%v\n", err2)
		return
	} else {
		fmt.Printf("写入的字节数是: %d\n", n)
	}
	// writestring
	n, _ = fileObj.WriteString("baby")
	fmt.Printf("写入的字节数是: %d\n", n)
	fileObj.Close()
}

func writeFile2() {
	// O_APPEND 追加  O_TRUNC 每次写都清空
	fileObj, err1 := os.OpenFile("./src/day05/10_file_write/test.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Printf("open file error :%v\n", err1)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	ret := bufio.NewWriter(fileObj)
	ret.WriteString("let's Go\n") // 写到缓存中
	ret.Flush()                   // 将缓存中的内容写入到文件
}

func writeFile3() {
	str := "simiada"
	err := ioutil.WriteFile("./src/day05/10_file_write/test.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write error: %v\n", err)
	} else {
		fmt.Println("write done")
	}
}

func main() {
	//writeFile1()
	//writeFile2()
	writeFile3()
}
