package main

import "fmt"

// 结构体模拟实现其他语言中的“继承”
// 动物的结构体
type animal struct {
	name string
}

// 动物的方法
func (a animal) move() {
	fmt.Printf("%s is moving \n", a.name)
}

// 狗结构体类
type dog struct {
	dogfeet uint
	animal  //此dog结构体包含了animal结构体， animal拥有的方法，dog结构体也能调用
}

// 狗的方法
func (d dog) shout() {
	fmt.Printf("%s is shouting \n", d.animal.name)
}

func main() {
	d1 := dog{
		dogfeet: 4,
		animal: animal{
			name: "dog",
		},
	}
	d1.move()
	d1.shout()
	fmt.Println(d1)
}
