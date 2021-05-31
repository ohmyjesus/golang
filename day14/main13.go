package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 重定向 http://localhost:8080/redirect 到  http://www.baidu.com

func main() {
	// 创建路由
	r := gin.Default()
	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 默认端口8080
	r.Run()

}
