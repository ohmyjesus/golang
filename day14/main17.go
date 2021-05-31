package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// cookie的使用

func main() {
	// 创建路由
	r := gin.Default()
	// 注册中间件
	r.GET("cookie", func(c *gin.Context) {
		// 判断客户端是否携带了cookie
		str, err := c.Cookie("key_cookie")
		if err != nil {
			// 说明没有携带cookie 应该设置cookie
			str = "NotSet"
			// SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
			// maxAge int  cookie维持的时间 单位秒
			// path string cookie所在目录
			// domain string 域名
			// secure bool 是否只能通过https访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie的值是%s\n", str)
	})

	// 默认端口8080
	r.Run(":8000")
}
