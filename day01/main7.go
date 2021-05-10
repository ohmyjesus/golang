package main

import "fmt"

// for循环
// 9x9 乘法表
func helper() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", i, j, j*i)
		}
		fmt.Println()
	}
}

func main() {
	// 基本格式  多用这种
	count := 2
	for i := 0; i < count; i++ {
		fmt.Printf("%d\n", i)
	}

	// 变种1  省略初始条件
	var i int = 5
	for ; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// 变种2   省略初始条件和末尾条件
	i = 5
	for i < 10 {
		fmt.Printf("%d\n", i)
		i++
	}

	// for range 遍历字符串  返回索引和值
	s := "hello"
	for i, v := range s {
		fmt.Printf("%d, %c\n", i, v)
	}

	helper()
	// 对于英文字符为byte， 对于中文和其他类型称为rune

	// 字符串修改  字符串是不能修改的，需要强转为 []byte 或 []rune ，完成后再转化为string类型
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换为rune切片  ['白'，'萝'，'卜']
	s3[0] = '红'      //是由一个个字符构成的 需要为字符类型
	fmt.Println(string(s3))

	c3 := "H" //string
	c4 := 'H' //int32
	fmt.Printf("c3:%T c4:%T\n", c3, c4)

	// 类型转换
	n := 10
	var f float32
	f = float32(n)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
