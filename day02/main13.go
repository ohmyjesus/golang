package main

import "fmt"

// map

func main() {
	// 声明的时候必须初始化
	var v1 = make(map[string]int, 10)
	v1["x"] = 1
	v1["y"] = 2
	fmt.Println(v1)

	// 约定俗成，用ok接收返回的布尔值
	fmt.Println(v1["z"]) //如果不存在这个key，则拿到0值
	value, ok := v1["z"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no this key")
	}

	//map的遍历 for range
	for k, v := range v1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range v1 {
		fmt.Println(k)
	}

	// 只遍历value
	for _, v := range v1 {
		fmt.Println(v)
	}

	// 删除map的元素
	delete(v1, "y")
	fmt.Println(v1)
	delete(v1, "z")
	fmt.Println(v1)
}
