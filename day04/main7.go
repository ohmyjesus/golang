package main

import "fmt"

// 方法和接收者

// Go语言中如果标识符首字母是大写的，就表示对外部是可见的，是公有的

// 狗的结构体
type dog struct {
	name string
}

// 狗的构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是一种作用于特定类型变量的函数
// 方法又分为值类型的接收者和指针类型的接收者
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {}
func (d dog) wang() {
	fmt.Printf("%s 旺旺\n", d.name)
}

// 人的结构体
type person struct {
	name string
	age  int
}

// 人的构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 人的方法
// 指针类型的接收者
func (p *person) guonian() {
	p.age++
}

func main() {
	d1 := newDog("小明")
	d1.wang()
	p1 := newPerson("xyh", 18)
	p1.guonian()
	fmt.Println(p1.age)
}
