package main

import "fmt"

// goto语句 可以实现跳出多层循环

func main() {
	count := 3
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if j == 1 {
				goto label
			}
			fmt.Printf("%d - %d\n", i, j)
		}
	}

label: //标签
	fmt.Println("over")
}
