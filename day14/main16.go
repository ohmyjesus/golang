package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件练习 执行顺序为 先中间件 -- 自己的函数 -- 中间件后续的事情
// 定义程序计时中间件，然后定义两个路由，执行函数结束后应该打印统计的计时时间

// 定义中间件
func myTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情
		// 打印程序执行时间
		fmt.Println("程序用时: ", time.Since(start))

	}
}

func shoppingIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(time.Second * 3)
	}
}

func shoppingHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(time.Second * 5)
	}
}

func main() {
	// 创建路由
	r := gin.Default()
	// 注册中间件
	r.Use(myTime())
	// {}是为什么代码规范
	// 定义路由组
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shoppingIndex())
		shoppingGroup.GET("/home", shoppingHome())
	}

	// 默认端口8080
	r.Run(":8000")
}
