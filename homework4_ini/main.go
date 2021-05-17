package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig  mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig  redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

// 读文件
func loadIni(fileName string, data interface{}) (err error) {
	//0. 参数的校验
	//0.1 传进来的data参数必须是指针类型，因为需要在函数中对其赋值
	t := reflect.TypeOf(data)
	fmt.Println(t.Kind())
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data param is not a ptr") //和fmt.Sprintf用法一样  不过是一个返回字符串类型，一个返回error类型
		return
	}

	//0.2 传进来的data参数必须是结构体类型指针，因为配置文件中有各种键值对，需要赋值给结构体字段
	fmt.Println(t.Elem().Kind())
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data param is not a struct") //和fmt.Sprintf用法一样  不过是一个返回字符串类型，一个返回error类型
		return
	}

	//1. 读文件得到字节类型的数据
	var b []byte
	b, err = ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("open file wrong")
		return
	}
	s := strings.Split(string(b), "\r\n") // 按照换行符来进行分割  返回值类型为string类型的切片，每个元素就是一行的内容
	fmt.Println(s[1])

	//2. 一行一行地从s中拿出来

	var structName string
	for idx, line := range s {
		var s1 string = ""
		// 如果为空行则直接跳过
		if s[idx] == "" {
			continue
		}

		//2.1 如果是注释，就跳过
		if line[0] == '#' || line[0] == ';' {
			continue
		}
		//2.2 如果是[] 就表示是节section
		if line[0] == '[' {
			for i := 1; line[i] != ']'; i++ {
				s1 += string(line[i])
			}

			// 拿到了节section根据反射去找对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ { // 因为t是指针，所以需要加elem() 否则直接t.NumField()
				field := t.Elem().Field(i)
				if field.Tag.Get("ini") == s1 { // mysql redis
					// 说明找到了对应结构体，记录当前节名mysql对应的字段名MysqlConfig
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", s1, structName)
				}
			}
			// 开始下次循环
			continue
		}

		//2.3 如果不是[] 就是=分割的键值对
		// 分割=左右两边  等号左边是key，右边是value
		index := strings.Index(line, "=")
		key := line[:index]
		value := line[index+1:]

		// 根据structname 去 data 里面取出结构体
		v := reflect.ValueOf(data)
		sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的类型信息
		sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

		if sType.Kind() != reflect.Struct {
			err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
		}

		var structKey string
		var fileType reflect.StructField
		// 3. 遍历嵌套结构体中的每一个字段，判断tag是不是等于key
		for i := 0; i < sValue.NumField(); i++ {
			field := sType.Field(i) //tag类型是存放在type类型信息中的
			fileType = field
			if field.Tag.Get("ini") == key {
				// 找到对应的字段
				structKey = field.Name
				break
			}
		}

		// 4. 如果等于，就给tag赋值
		// 根据structKey去mysql或者redis取出字段赋值
		temp := sValue.FieldByName(structKey) // temp = Address / Port ...
		fmt.Println(structName, fileType.Type.Kind())
		switch fileType.Type.Kind() {
		case reflect.String:
			temp.SetString(value)
		case reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int:
			var i int64
			i, err = strconv.ParseInt(value, 10, 64) // 将字符串类型的数字 转为 数字   (s, 十进制, 64位)
			if err != nil {
				err = fmt.Errorf("line: %d value type error", idx+1)
				return
			}
			temp.SetInt(i)
		}
	}
	return
}

func main() {
	// 类型信息用TypeOf
	// 值信息用ValueOf
	var cfd Config
	err := loadIni("F:/GoCode/src/day06/07_ini/conf.ini", &cfd)
	if err != nil {
		fmt.Printf("load ini failed :%v\n", err)
		return
	}
	fmt.Println("---")
	fmt.Println(cfd.MysqlConfig.Address, cfd.MysqlConfig.Port, cfd.MysqlConfig.Username, cfd.MysqlConfig.Password)
	fmt.Println(cfd.RedisConfig.Host, cfd.RedisConfig.Port, cfd.RedisConfig.Password, cfd.RedisConfig.Database)
}
