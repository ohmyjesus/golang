package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// 上传单个文件

// 先启动后台程序，再打开html代码open in default browser，输入内容之后点击登录即可

func main() {
	// 1. 创建路由
	r := gin.Default()
	// 2. api参数
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("from file error:", err)
			return
		}
		log.Println(file.Filename) // log默认打印时间信息
		// 将文件传到同级目录中
		c.SaveUploadedFile(file, file.Filename)
		c.String(200, fmt.Sprintf("%s upload!\n", file.Filename))
	})

	// 3. 监听端口  默认8080端口  也可自定义端口
	r.Run()
}
