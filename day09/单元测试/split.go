package split

import (
	"strings"
)

// 自定义函数 切割字符串
// "abce", "b" --> [a ce]
func Sp(str string, sep string) []string {
	var ret []string = make([]string, 0, strings.Count(str, sep)+1) // 优化代码，只申请一次内存
	n := strings.Index(str, sep)
	length := len(sep)
	for n >= 0 {
		ret = append(ret, str[:n]) // 把一个字符串类型追加到字符串的切片里
		str = str[n+length:]
		n = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
