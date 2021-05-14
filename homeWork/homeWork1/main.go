package main

import (
	"fmt"
	"os"
)

/*
	方法版学生管理系统
	写一个系统能够查看/新增/删除学生
*/

//管理员的结构体
type manage struct {
	num      int
	password int
}

// 学生的结构体
type student struct {
	name string
	id   int
}

// map映射
var allstu = make(map[int]*student, 20)

// 构造函数
func newStu(name string, id int) *student {
	return &student{
		name: name,
		id:   id,
	}
}

// 查看学生的函数
func (manage) lookStu() {
	fmt.Println("学生信息如下：")
	// 遍历打印
	for _, v := range allstu {
		fmt.Printf("学生姓名: %s\t学号: %d\n", v.name, v.id)
	}
	fmt.Println()
}

// 新增学生的函数
func (manage) addStu() {
	var (
		name string
		id   int
	)
	fmt.Print("请输入添加学生的学号: ")
	fmt.Scanln(&id)
	// 判断学号是否存在
	for _, v := range allstu {
		if id == v.id {
			fmt.Println("此学号已存在，添加失败")
			return
		}
	}
	fmt.Print("请输入添加学生的姓名: ")
	fmt.Scanln(&name)
	ret := newStu(name, id)
	allstu[id] = ret
	fmt.Println(" ---- 添加成功 ---- ")
}

// 删除学生的函数
func (manage) delStu() {
	fmt.Print("请输入你想删除学生的学号: ")
	var id int
	fmt.Scanln(&id)
	// 根据键删除元素
	for _, v := range allstu {
		if id == v.id {
			delete(allstu, id) // 根据键进行删除
			fmt.Println("删除成功")
			fmt.Println()
			return
		}
	}
	fmt.Println("此学号不存在，删除失败")
	fmt.Println()
}

func login() {
	var (
		zhanghao int
		mima     int
	)
	for {
		fmt.Print("请输入管理员账号(提示123): ")
		fmt.Scanln(&zhanghao)
		fmt.Print("请输入管理员密码(提示123): ")
		fmt.Scanln(&mima)
		fmt.Println()
		if zhanghao == 123 && mima == 123 {
			user.num = zhanghao
			user.password = mima
			return
		} else {
			fmt.Println("账号密码不对，请重新输入: ")
		}
	}

}

var user manage

func main() {
	login()
	fmt.Println(" --- 欢迎光临学生系统 --- ")
	for {
		// 1. 打印菜单
		fmt.Println(`   
				1. 查看学生   
				2. 新增学生   
				3. 删除学生
				4. 退出
		`)
		fmt.Print("请输入你的选择: ")
		var choice int
		fmt.Scanf("%d\n", &choice)
		// 2. 用户选择并执行对应的函数
		switch choice {
		case 1:
			user.lookStu()
		case 2:
			user.addStu()
		case 3:
			user.delStu()
		case 4:
			fmt.Println("您已安全退出~")
			os.Exit(1)
		}
	}
}
