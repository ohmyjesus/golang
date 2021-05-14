package main

import (
	"fmt"
	"io"
	"os"
)

// 完成：向文件中间插入数据, 需要利用到一个临时文件
// 为了突出逻辑，我并未对文件操作函数做错误判断

func openFile() {
	// 打开文件
	fileObj, err := os.OpenFile("F:/GoCode/src/day06/01_file_insert/test.txt", os.O_RDWR, 0644)
	if err == nil {
		fmt.Println("开始操作源文件")
	}
	defer fileObj.Close()

	// 指针量偏移到要插入的源文件中的位置
	fileObj.Seek(1, 0)

	// 打开临时文件
	filetemp, _ := os.OpenFile("F:/GoCode/src/day06/01_file_insert/temp.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644) //TRUNC 打开时清空临时文件

	// 将原文件读到的数据循环写入到临时文件中
	var s [1]byte
	for {
		_, err := fileObj.Read(s[:]) //读到s中
		if err == io.EOF {           // 表示读到了文件末尾
			break
		}
		filetemp.Write(s[:]) //写到临时文件中
	}

	// 源文件指针重新偏移
	fileObj.Seek(1, 0)

	// 要插入到原文件中的内容
	fileObj.WriteString("\nb")

	// 临时文件指针重新偏移
	filetemp.Seek(0, 0)

	// 再把临时文件的内容循环写回到原文件中
	var s2 [1]byte
	for {
		_, err := filetemp.Read(s2[:])
		if err == io.EOF {
			break
		}
		fileObj.Write(s2[:])
	}

	fmt.Println("执行完成")
}

func main() {
	openFile()
}
