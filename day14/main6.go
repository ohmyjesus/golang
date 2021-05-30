package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传多个文件

// 先启动后台程序，再打开html代码open in default browser，输入内容之后点击登录即可

func main() {
	// 1. 创建路由
	r := gin.Default()
	// 限制表单上传大小8MB  默认上传大小32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s\n", err))
			return
		}
		// 获取所有文件
		files := form.File["files"]
		// 遍历所有图片
		for _, v := range files {
			// 逐个存
			err = c.SaveUploadedFile(v, v.Filename)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("save file err %s\n", err))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d file\n", len(files))) // 加载了多少个图片
	})

	// 3. 监听端口  默认8080端口  也可自定义端口
	r.Run()
}
