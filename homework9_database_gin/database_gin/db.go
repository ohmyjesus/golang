package templates

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() (err error) {
	// 数据库信息
	// 用户名：密码@tcp(ip:port)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/databasegin" // 不会校验用户名和密码是否正确
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10) // 如果这10个连接数都被占用了，那么下一个连接会阻塞等待其他连接释放
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

// 查询多条数据
func SelectMoreRow() (b []Book, err error) {
	sqlStr2 := `select * from book`
	// 需要传地址进去
	err = db.Select(&b, sqlStr2)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	return
}

// 插入数据
func InsertRow(title string, price int64) (err error) {
	// 1. sql语句
	sqlStr := `insert into book (title, price) values(?, ?)`
	// 2. exec执行
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入失败:", err)
		return
	}
	return
}

// 删除数据
func DeleteRow(id int64) (err error) {
	sqlStr := `delete from book where id = ?`
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败:", err)
		return
	}
	return
}
