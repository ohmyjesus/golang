package main

import "fmt"

// if 条件判断

func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("已成年")
	// } else {
	// 	fmt.Println("未成年")
	// }

	// 作用域 age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
	//fmt.Println(age) 此时age是不存在的
}
