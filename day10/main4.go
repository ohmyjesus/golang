package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 不导入这个包则报错 sql: unknown driver "mysql" (forgotten import?)
	"github.com/jmoiron/sqlx"
)

// sqlx 例子
// 查询时候比普通的查询更方便一点

var db *sqlx.DB

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

// 查询单条数据
func selectRow() {
	sqlStr1 := `select id, name, age from user where id = 1`
	var u1 user
	// get方法查询
	// 第一需要地址进去，第二是get方法它是利用反射进行，为了使其他包知道传进来的结构体有哪些数据，结构体属性的首字母需要大写，对其他包可见
	db.Get(&u1, sqlStr1)
	fmt.Printf("%#v\n", u1)
}

// 查询多条数据
func selectMoreRow() {
	var u2 []user = make([]user, 0, 10) // slice, map, channel --> make
	sqlStr2 := `select * from user`
	// 需要传地址进去
	db.Select(&u2, sqlStr2)
	fmt.Printf("u2: %#v\n", u2)
}

func main() {
	err := initMysql()
	if err != nil {
		fmt.Println("initmysql error", err)
		return
	}
	selectRow()
	selectMoreRow()
}
