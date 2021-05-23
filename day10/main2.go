package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 连接mysql

var db *sql.DB

// 用户结构体
type user struct {
	name string
	id   int
	age  int
}

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

// 查询单个记录
func queryOne(id int) {
	var u1 user
	// 1. 查询单条的sql语句
	sqlStr := "select id, name, age from user where id=?;" // ?是占位符
	// 2. 执行
	for i := 0; i < 11; i++ {
		db.QueryRow(sqlStr, id) // 从连接池拿一个连接出来，去数据库查询单条数据记录
		fmt.Printf("已经进行了第%d次查询\n", i)
	}
	// 3. 拿到结果
	//rowObj.Scan(&u1.id, &u1.name, &u1.age) // 必须对rowObj调用scan方法，因为它会释放数据库连接，且需要传一个地址进去
	// 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多条记录
func queryMore(id int) {
	// 1. sql语句
	sqlStr := "select id, name, age from user where id>?;" // ?是占位符
	// 2. 执行
	rowsObj, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Println("query error ", err)
		return
	}
	// 3. 一定要手动关闭，延时释放数据库连接
	defer rowsObj.Close()
	// 4. 循环取值 使用内置方法next
	for rowsObj.Next() {
		var u2 user
		rowsObj.Scan(&u2.id, &u2.name, &u2.age) // 和上面的scan方法不一样，上面的scan自带关闭，这里需要手动关闭
		fmt.Printf("u2:%#v\n", u2)
	}

}

// 插入数据
func insertRow() {
	// 1. sql语句
	sqlStr := `insert into user (name, age) values("xyh", 18)`
	// 2. exec执行
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("insert exec error", err)
		return
	}
	// 如果是插入数据的操作，则能够拿到插入数据的id值
	newID, _ := res.LastInsertId()
	fmt.Printf("在第%d行插入了数据\n", newID)
}

// 更新数据
func updateRow(newAge int, id int) {
	sqlStr := `update user set age = ? where id = ?` // sql语句是可以自定义的
	res, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Println("update exec error ", err)
		return
	}
	// RowsAffected返回被影响(更新)了几行数据
	n, _ := res.RowsAffected()
	fmt.Printf("更新了%d行数据\n", n)

}

// 删除数据
func deleteRow(id int) {
	sqlStr := `delete from user where id = ?`
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete exec error", err)
		return
	}
	// RowsAffected返回被影响(删除)了几行数据
	n, _ := res.RowsAffected()
	fmt.Printf("删除了\n行数据", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user (name, age) values (?, ?)`
	stmt, err := db.Prepare(sqlStr) // 把sql语句先发给mysql预处理一下
	if err != nil {
		fmt.Println("prepare error", err)
		return
	}
	defer stmt.Close()
	// 后续只需要拿到stmt去执行一些操作
	var m map[string]int = make(map[string]int, 10) // map存储
	m = map[string]int{
		"kong1": 23,
		"kong2": 24,
	}
	// 插入数据
	for k, v := range m {
		stmt.Exec(k, v)
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("initDB error", err)
		return
	}
	fmt.Println("数据库连接成功")
	queryOne(2)
	queryMore(0)
	insertRow()
	updateRow(168, 3)
	deleteRow(4)
	prepareInsert()
}
