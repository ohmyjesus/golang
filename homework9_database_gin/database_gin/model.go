package templates

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    int64  `form:"id" json:"id" uri:"id" xml:"id" `
	Title string `form:"title" json:"title" uri:"title" xml:"title"`
	Price int64  `form:"price" json:"price" uri:"price" xml:"price"`
}

var r *gin.Engine

func InitGin() {
	// 创建路由
	r = gin.Default()

	// 加载模板文件
	r.LoadHTMLGlob("./database_gin/*") // 路径

	// 创建路由组
	v1 := r.Group("/book")
	v1.GET("/list", BookListHandler)                       // http://localhost:8000/book/list
	v1.GET("/new", BookInsertHandler)                      // http://localhost:8000/book/new
	v1.GET("/delete", MiddleWare(), func(c *gin.Context) { // http://localhost:8000/book/delete
		// 重定向到主页面
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/book/list")
	})

	// 默认端口8080
	r.Run(":8000")
}

// 主页面
func BookListHandler(c *gin.Context) {
	b, err := SelectMoreRow()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err})
		return
	}
	// HTML模板渲染
	c.HTML(200, "book_list.html", gin.H{"code": 0, "data": b})
}

// 添加页面
func BookInsertHandler(c *gin.Context) {
	// 渲染HTML
	c.HTML(200, "new_book.html", gin.H{})
	// 获取表单数据
	var book Book // // 声明接收的变量, 结构体类型
	r.POST("book/new", func(c *gin.Context) {
		// bind()  默认解析并绑定form格式
		// 根据请求中conteny-type自动推断
		err := c.Bind(&book)
		if err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = InsertRow(book.Title, book.Price)
		if err != nil {
			fmt.Println("插入数据失败：", err)
			return
		}
		// 重定向到主页面
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/book/list")
	})

}

// 中间件  获取URL参数
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// http://localhost:8080/book/delete?id=5 --> 5 指定URL参数则显示该参数
		str := c.DefaultQuery("id", "1") // http://localhost:8080//book/delete --> 1 直接访问则显示默认值
		// 字符串整数转化为整型整数
		id, _ := strconv.Atoi(str)
		err := DeleteRow(int64(id))
		if err != nil {
			fmt.Println("删除失败", err)
			return
		}
		fmt.Printf("删除成功 id:%d\n", id)
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情

	}
}
