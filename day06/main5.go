package main

import (
	"encoding/json"
	"fmt"
)

// json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"xyh","age":9000}`
	var p person
	json.Unmarshal([]byte(str), &p) // 映射  运行的时候才确定类型
	fmt.Println(p.Name, p.Age)
}
