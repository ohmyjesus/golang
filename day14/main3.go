package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 取出gin的URL参数 ?后面是URL参数比如 name=xyh

func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) { // 必须要加:和*号
		// DefaultQuery的第二个参数是默认值
		str := c.DefaultQuery("name", "jack") // http://localhost:8080/welcome --> jack 直接访问显示默认值
		// http://localhost:8080/welcome?name=xyh --> xyh 指定URL参数则显示该参数
		c.String(http.StatusOK, str)
	})
	r.Run()
}
