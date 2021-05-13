package main

import "fmt"

// 同一个结构体实现多个接口
// 接口嵌套

type animal interface {
	mover
	eater
}
type mover interface {
	move()
}

type eater interface {
	eat(string)
}

// dog实现了mover接口和eater接口
type dog struct {
	name string
}

func (d *dog) move() {
	fmt.Println("dog is moving")
}
func (d *dog) eat(something string) {
	fmt.Printf("dog is eating %s\n", something)
}

func main() {

}
