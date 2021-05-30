package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin的基本路由 此时访问网页 localhost:8080 即可显示字符串 hello world!

func f1(c *gin.Context) {
	fmt.Println("hahha")
}

func main() {
	// 1. 创建路由
	// 默认使用了两个中间件 Logger(), Recovery() 也可以创建不带中间件的路由 r:=gin.new()
	r := gin.Default()
	// 2. api参数 和 绑定路由规则，执行的函数
	// gin.context 封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!") // 在网页上输出hello world
	})
	r.POST("xxx/post", f1)
	r.PUT("xxx/put")
	// 3. 监听端口  默认8080端口  也可自定义端口
	r.Run()
}
