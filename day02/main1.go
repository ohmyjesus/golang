package main

import "fmt"

// for循环之跳出循环

func main() {
	// 当i = 5 时就跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// 当i = 5 时就跳过此for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
