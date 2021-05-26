package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 专门向kafka写日志的模块

var (
	client sarama.SyncProducer // 声明一个全局的连接kafka的生产者client
)

// 初始化操作 -- 利用sarama向kafka里写日志
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀一个partition分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	return
}

// 向kafka发送一个消息
func SendToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送到kafka中
	pid, offset, err := client.SendMessage(msg)

	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	fmt.Println("发送成功")
}
