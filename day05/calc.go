package calc

import "fmt"

// 包中的标识符如果是小写的则表示私有的
// 首字母大写的才能被别的包调用

func init() {
	fmt.Println("this is package calc init method")
}

func Add(x, y int) int {
	return x + y
}
