package main

import (
	"database/sql" //原生支持连接池，并发安全
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 连接mysql

func main() {
	// 数据库信息
	// 用户名：密码@tcp(ip:port)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/goday10" // 不会校验用户名和密码是否正确
	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql open error ", err) // dsn的格式不正确时会报错
		return
	}
	err = db.Ping() // 验证用户名和密码 尝试连接数据库
	if err != nil {
		fmt.Println("ping error ", err)
	}
	fmt.Println("数据库连接成功")
}
