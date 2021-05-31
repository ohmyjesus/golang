package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局中间件和局部中间件   当访问页面时执行顺序为 先中间件 -- 自己的函数 -- 中间件后续的事情

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("中间件开始执行")
		// 设置变量到context的key中,可以通过get取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		// 程序执行时间
		fmt.Println(time.Since(start))
		//fmt.Println(time.Now().Sub(start))

	}
}

func main() {
	// 创建路由
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	// {}是为什么代码规范
	{
		r.GET("/middleware", func(c *gin.Context) { // 访问http://localhost:8080/middleware 只走一次全局中间件
			// 取值
			v, exists := c.Get("request")
			if !exists {
				fmt.Println("key 不存在")
				return
			}
			fmt.Println("request: ", v)
			// 页面返回
			c.JSON(200, gin.H{"request": v})
		})
		// 局部中间件
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) { // 访问http://localhost:8080/middleware2 先走全局中间件，再走自己的中间件
			// 取值
			v, exists := c.Get("request")
			if !exists {
				fmt.Println("key 不存在")
				return
			}
			fmt.Println("request: ", v)
			// 页面返回
			c.JSON(200, gin.H{"request": v})
		})
	}

	// 默认端口8080
	r.Run()

}
