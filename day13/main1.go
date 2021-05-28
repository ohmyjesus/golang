package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

// ES 插入数据  执行此程序之前需要启动elasticsearch
// 路径如下 F:\elasticsearch\elasticsearch-7.13.0\bin>elasticsearch.bat

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Person{Name: "rion", Age: 22, Married: false} // 向ES中插入数据，可以通过postman查询到此条数据
	put1, err := client.Index().
		Index("user").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type) // Indexed user ACXtsHkBRSaraA_4ox49 to index user, type go
}
