package main

import (
	"github.com/gin-gonic/gin"
)

// HTML模板渲染

func main() {
	// 创建路由
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("index.tmpl") // 路径
	// r.LoadHTMLFiles("/index.tmpl") // 文件名
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终json将title替换
		c.HTML(200, "index.tmpl", gin.H{"title": "我的标题"})
	})

	// 默认端口8080
	r.Run()

}
