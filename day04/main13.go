package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
// 1. 序列化 ： Go语言中的结构体变量 --> json格式的字符串
// 2. 反格式化：json格式的字符串 --> Go语言中的结构体变量

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "熊宇航",
		Age:  18,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	// 反序列化
	str := `{"name":"熊宇航","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
