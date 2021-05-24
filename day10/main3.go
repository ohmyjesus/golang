package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// mysql事务

var db *sql.DB

// 初始化数据库
func initDB() (err error) {
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

func transAction() {
	// 开启事务  事务的原子性，不可分割
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("db begin error", err)
		return
	}
	// 执行多个sql操作 将id为6的年龄减2岁，id为7的年龄加2岁
	sqlStr1 := `update user set age = age - 2 where id = 6`
	sqlStr2 := `update xxx set age = age + 2 where id = 7`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 如果执行失败则回滚
		tx.Rollback()
		fmt.Println("第一条sql语句执行失败")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 如果执行失败则回滚
		tx.Rollback()
		fmt.Println("第二条sql语句执行失败")
		return
	}
	// 上面两步sql语句都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("commit error", err)
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("initDB error", err)
		return
	}
	fmt.Println("数据库连接成功")
	transAction()
}
