package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 不导入这个包则报错 sql: unknown driver "mysql" (forgotten import?)
)

// sql注入

var db *sql.DB

// 用户结构体
type user struct {
	Name string // 首字母大写，别的包才能调用到
	Id   int
	Age  int
}

// 初始化数据库
func initMysql() (err error) {
	// 数据库信息
	// 用户名：密码@tcp(ip:port)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test" // 不会校验用户名和密码是否正确
	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() // 验证用户名和密码 尝试连接数据库
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10) // 如果这10个连接数都被占用了，那么下一个连接会阻塞等待其他连接释放
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

// mysql注入 在拼接字符串的时候搞小动作   sql注入解决办法：不要拼接字符串或通过预处理方式
// 单条查询
func sqlInjectDemo(name string) {
	// 自己拼接一个sql语句的字符串
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name) // select id, name, age from user where name='dd' or 1=1 #' where条件永远成立 #注释
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

func main() {
	initMysql()
	// sqlInjectDemo("xyh")
	sqlInjectDemo("dd' or 1=1 #")
	sqlInjectDemo("dd' union select * from user #")
}
