package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin数据解析和绑定

// 2. 表单数据解析和绑定到对象上

type Login struct {
	// binding:"required" 修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()
	// json绑定
	r.POST("loginForm", func(c *gin.Context) { // 这里的loginForm和HTML里面的字段相呼应
		// 声明接收的变量, 结构体类型
		var form Login
		// bind()  默认解析并绑定form格式
		// 根据请求中conteny-type自动推断
		err := c.Bind(&form)
		if err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名和密码是否正确
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			fmt.Println("---", form.Password, form.User)
			return
		}
		fmt.Println("---", form.Password, form.User)
		c.JSON(http.StatusOK, gin.H{"status": "200"})

	})

	// 默认端口8080
	r.Run()
}
