package main

import "fmt"

// 值接收者和指针接收者实现接口的区别
// 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存

// 使用指针接收者实现接口，只能存结构体指针类型的变量。
type animal interface {
	move()
	eat(string)
}

type dog struct {
	name string
}

// 使用值接收者实现接口的方法
// func (d dog) move() {
// 	fmt.Println("dog is moving")
// }
// func (d dog) eat(something string) {
// 	fmt.Printf("dog is eating %s\n", something)
// }

// 使用指针接收者实现接口的方法
func (d *dog) move() {
	fmt.Println("dog is moving")
}
func (d *dog) eat(something string) {
	fmt.Printf("dog is eating %s\n", something)
}

func main() {
	var a1 animal
	var d1 = dog{ // dog
		"tom",
	}
	var d2 = &dog{ // *dog
		"Tom",
	}
	a1 = &d1 // 实现animal这个接口的是dog的指针类型
	a1 = d2
	fmt.Println(a1)

}
