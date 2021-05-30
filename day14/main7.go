package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// routesGroup路由组

func login(c *gin.Context) {
	// 设置默认值
	name := c.DefaultQuery("name", "xyh")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "ly")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func main() {
	// 创建路由
	r := gin.Default()

	// 路由组1 处理get请求
	v1 := r.Group("/v1")
	v1.GET("/login", login)  // 访问http://127.0.0.1:8080/v1/login?name=xxx  -->  hello xxx
	v1.GET("submit", submit) // 访问http://127.0.0.1:8080/v1/submit?name=xx  -->  hello xx

	// 路由组2
	v2 := r.Group("/v2")
	v2.POST("/login", login)
	v2.POST("submit", submit) // 终端模式下>curl http://127.0.0.1:8080/v2/submit -X POST --> hello ly
	// 3. 监听端口  默认8080端口  也可自定义端口
	r.Run()
}
