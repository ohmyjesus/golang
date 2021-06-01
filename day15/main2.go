package main

import "fmt"

// 选项设计模式

type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// 声明一个函数类型的变量，用于传参
type Option func(opts *Options)

// 初始化结构体方式1 -- 写死的
func InitOptions1(strOption1, strOption2, strOption3 string, intOption1, intOption2, intOption3 int) *Options {
	return &Options{
		strOption1: strOption1,
		strOption2: strOption2,
		strOption3: strOption3,
		intOption1: intOption1,
		intOption2: intOption2,
		intOption3: intOption3,
	}
}

// 初始化结构体方式2 -- 还是写死的
func InitOptions2(options ...interface{}) *Options {
	option := Options{}
	// 遍历参数
	for index, v := range options {
		switch index {
		case 0:
			str, ok := v.(string)
			if !ok {
				return &option
			} else {
				option.strOption1 = str
			}
		}
	}
	return &option
}

// 引入选项设计模式来解决此问题
func InitOptions3(opts ...Option) {
	options := &Options{}
	// 遍历opts，得到每一个函数
	for _, opt := range opts {
		// 调用函数，在函数里，给传进去的对象赋值
		opt(options)
	}
	fmt.Printf("%#v\n", options)
}

// 定义具体给某个字段赋值的方法
func WithStrOption1(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}

func main() {
	InitOptions3(WithStrOption1("str1"))
}
