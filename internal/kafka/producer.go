package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"go-practics/config"
)

// RunProducer 运行Kafaka生产者
func RunProducer() {
	kURL := fmt.Sprintf("%s:%d", config.Host, config.KafkaPort)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kURL})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// 生成消息的传递报告处理程序
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// 向主题生成消息（异步）
	topic := config.KafkaTopic
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// 等待所有消息被发送
	p.Flush(15 * 1000)
	fmt.Println("Message sent successfully")
}
