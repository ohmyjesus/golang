package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// net.http  server
// ResponseWriter服务器端响应     Request客户端请求
func f1(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("F:/GoCode/src/day09/02_http_server/test.html")
	if err != nil {
		fmt.Println("readfile error ", err)
	}
	// 服务器端将b写到响应中，返回给客户端，浏览器将数据b按照HTML/CSS的方式来渲染页面
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在URL上(query param)，请求体中是没有数据的，所以body为空[]
	queryParam := r.URL.Query()         // 自动帮我们识别URL中的query param
	name := queryParam.Get("name")      // 123
	age := queryParam.Get("age")        // 456
	fmt.Println(name, age)              // 123 456
	fmt.Println(r.URL)                  // /ly?name=123&age=456
	fmt.Println(r.Method)               // GET
	fmt.Println(ioutil.ReadAll(r.Body)) // [] <nil> 在服务端打印客户端发来的请求
	w.Write([]byte("ok"))               // 服务器端的响应，数据写回给客户端
}

func main() {
	//http://192.168.123.176:8080/xyh
	http.HandleFunc("/xyh", f1) //访问/xyh就执行f1函数
	http.HandleFunc("/ly", f2)  //访问/ly就执行f2函数
	http.ListenAndServe("0.0.0.0:8080", nil)
}
