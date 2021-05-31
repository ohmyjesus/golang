package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// cookie的练习题
// 实现功能：
// 1. 有两个路由，login和home
// 2. login用于设置cookie
// 3. home是访问查看信息的请求
// 4. 在请求home之前，先跑中间件代码，检验是否存在cookie

// 访问home，会显示错误，因为权限检验未通过
// 然后访问登录的请求，登录并设置cookie
// 再次访问home，访问成功

// 定义中间件1
func home(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断客户端是否携带了cookie
		_, err := c.Cookie("root")
		if err != nil {
			// 说明没有携带cookie 则提示禁止访问
			c.JSON(http.StatusUnauthorized, gin.H{"error": "statusUnauthorized"})
			// 同时后续不再执行
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": "home"})
		// 执行函数
		c.Next()
	}
}

// 定义中间件2
func login(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置cookie
		// SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
		// maxAge int  cookie维持的时间 单位秒
		// path string cookie所在目录
		// domain string 域名
		// secure bool 是否只能通过https访问
		// httpOnly bool 是否允许别人通过js获取自己的cookie
		c.SetCookie("root", "admin", 60, "/", "localhost", false, true)
		c.String(200, "login successful!")
		// 执行函数
		c.Next()
	}
}

func main() {
	// 创建路由
	r := gin.Default()
	// 定义路由组
	cookieGroup := r.Group("")
	{
		cookieGroup.GET("/login", login(r))
		cookieGroup.GET("/home", home(r))
	}

	// 默认端口8080  可指定端口
	r.Run(":8000")
}
