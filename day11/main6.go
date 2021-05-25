package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

// 利用tailf读日志

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开               文件到了一定大小，就关闭文件，打开新文件，相当于文件切割
		Follow:    true,                                 // 是否跟随               时间戳跟随上一个文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读  下一次启动时从上一次读到的位置继续读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n",
				tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
