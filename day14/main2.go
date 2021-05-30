package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 取出gin的API参数

// 单个匹配
func f1() {
	// 1. 创建路由
	r := gin.Default()
	// 取出api 参数   "/user/:name"为单个匹配  访问localhost:8080/user/xyh 显示xyh
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")       // 取出name的匹配参数
		c.String(http.StatusOK, name) // 在网页上输出该参数
	})
	r.Run()
}

// 任意匹配
func f2() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) { // 必须要加:和*号
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+action) // 访问localhost:8080/user/xyh/ly/xx 显示 xyh/ly/xx
	})
	r.Run()
}

func main() {
	f2()
}
