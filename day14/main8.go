package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin数据解析和绑定

// 1. json数据解析和绑定

type Login struct {
	// binding:"required" 修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()
	// json绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接收的变量, 结构体类型
		var json Login
		// 将request的body中的数据  按照json格式解析到结构体
		err := c.ShouldBindJSON(&json)
		if err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名和密码是否正确
		if json.User != "root" || json.Password != "admin" {
			// JSON 页面返回信息
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// 默认端口8080
	r.Run()
}
