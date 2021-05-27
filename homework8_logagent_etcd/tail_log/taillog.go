package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

// tailf读日志

// 初始化tailf
func Init(fileName string) (tails *tail.Tail, err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开               文件到了一定大小，就关闭文件，打开新文件，相当于文件切割
		Follow:    true,                                 // 是否跟随               时间戳跟随上一个文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读  下一次启动时从上一次读到的位置继续读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 //
	}
	tails, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	return
}

//  利用tailf读日志 利用通道传输消息 tails.Lines
func ReadChan(tails *tail.Tail) chan *tail.Line {
	return tails.Lines
}
