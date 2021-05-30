package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 表单参数

// 先启动后台程序，再打开html代码open in default browser，输入内容之后点击登录即可输出相应的字符串内容

func main() {
	// 1. 创建路由
	r := gin.Default()
	// 2. api参数
	r.POST("/form", func(c *gin.Context) {
		// 表单参数也可以设置默认值
		types := c.DefaultPostForm("type", "alert")
		// 接收非默认值
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 多选框
		hobbies := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobbies are %v", types, username, password, hobbies)) //拼接字符串
	})

	// 3. 监听端口  默认8080端口  也可自定义端口
	r.Run()
}
