package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka 消费者实例  此时就不要起终端打印了 因为是最新的偏移量offsetnewest

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("redis_log") // 注意这里的topic和你发消息的topic应该对应
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("分区列表", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("redis_log", int32(partition), sarama.OffsetNewest) // 最新的偏移量，不要再起终端打印
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}
	select {} //主goroutine一直在跑
}
