package main

// reflect

import (
	"fmt"
	"reflect"
)

type cat struct {
}

// 与类型断言类似   但是不是猜，而是直接获得类型
func reflectType(a interface{}) {
	ret := reflect.TypeOf(a)
	fmt.Printf("type: %v\n", ret)
	fmt.Printf("type name: %v  type kind: %v\n", ret.Name(), ret.Kind())
}

func main() {
	var c cat
	reflectType(c)
	var s string = "str"
	reflectType(s)
}
