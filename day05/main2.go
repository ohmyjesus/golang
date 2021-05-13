package main

import "fmt"

type dog struct {
	name string
}

type cat struct {
	name string
}

// dog的方法
func (d dog) move() {
	fmt.Println("dog is moving")
}
func (d dog) eat(something string) {
	fmt.Printf("dog is eating %s\n", something)
}

// cat的方法
func (c cat) move() {
	fmt.Println("cat is moving")
}
func (c cat) eat(something string) {
	fmt.Printf("cat is eating %s\n", something)
}

// 接口的实现
// 只要结构体实现了接口里面的方法，那么该结构体的类型也是此接口类型，就也能调用接口函数
type animal interface {
	move()
	eat(string)
}

// 接口函数
func f(a animal, something string) {
	a.eat(something)
	a.move()
}

func main() {
	var c1 cat
	c1.name = "mao"
	f(c1, "fish")
	var d1 dog
	d1.name = "gou"
	f(d1, "bune")
}
