package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 异步

func main() {
	// 创建路由
	r := gin.Default()
	// 1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		// 必须使用一个副本
		gc := c.Copy()
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println("异步执行: " + gc.Request.URL.Path)
		}()
	})

	// 2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		fmt.Println("同步执行: " + c.Request.URL.Path)
	})

	// 默认端口8080
	r.Run()

}
