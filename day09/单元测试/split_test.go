package split

// 1. 测试组和子测试
// 2. 性能基准测试

import (
	"reflect"
	"testing"
)

// // 测试组 go test -v   测试案例一起执行
// func TestSplit(t *testing.T) {
// 	// 将所有的信息放到结构体中
// 	type test struct {
// 		input string
// 		sep   string
// 		want  []string
// 	}
// 	// 结构体切片即结构体数组   每个元素都是一个结构体
// 	testGroup := []test{
// 		{input: "abcbef", sep: "b", want: []string{"a", "c", "ef"}},
// 		{input: "acef", sep: "b", want: []string{"acef"}},
// 		{input: "abcdef", sep: "bc", want: []string{"a", "def"}},
// 	}
// 	// 遍历结构体切片 前一个为索引，后一个为结构体本身
// 	for _, tc := range testGroup {
// 		got := Sp(tc.input, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) { // 是因为字符串不能直接比较，采用到了这个函数。 如果可以直接比较，则直接==比较
// 			// 测试用例失败时
// 			t.Errorf("wang %#v but got %#v", tc.want, got)
// 		}
// 	}
// }

// 子测试t.Run函数  go test -v  它可以单独跑某个测试案例 go test -run=TestSplit/case1
func TestSplit(t *testing.T) {
	// 将所有的信息放到结构体中
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义个map，其值为结构体类型
	var testGroup map[string]test = make(map[string]test, 1000) // 暂存1000个测试案例 具有自动扩容的能力
	testGroup = map[string]test{
		"case1": {input: "abcbef", sep: "b", want: []string{"a", "c", "ef"}},
		"case2": {input: "acef", sep: "b", want: []string{"acef"}},
		"case3": {input: "abcdef", sep: "bc", want: []string{"a", "def"}},
	}
	// 遍历结构体切片 前一个为索引，后一个为结构体本身
	for key, tc := range testGroup {
		t.Run(key, func(t *testing.T) {
			got := Sp(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				// 测试用例失败时
				t.Errorf("wang %#v but got %#v\n", tc.want, got)
			}
		})
	}
}

// BenchMark 的基准测试 go test -bench=Split -benchmem
func BenchmarkSp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sp("abbc", "bb")
	}
}

// 斐波那契数列的性能比较测试 go test -bench=Fib2
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
