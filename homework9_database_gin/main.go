package main

import (
	templates "day15/database_gin"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库作业
// 实现一个简易的图书馆管理系统，数据从mysql中读取，书籍有id，title，price三个字段

func main() {
	// 1. 连接数据库
	err := templates.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败 error: ", err)
		return
	}
	fmt.Println("连接数据库成功")

	// 2. 连接路由
	templates.InitGin()
 
}
