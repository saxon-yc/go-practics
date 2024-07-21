package kafka

import (
	"fmt"
	"go-practics/config"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func RunConsumer() {
	kUrl := fmt.Sprintf("%s:%d", config.Host, config.KafkaPort)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		/*
			bootstrap.servers
				这是一个必需的配置项，用于指定 Kafka 集群的地址。它可以是一个或多个 Kafka broker 的地址，
				格式为 host1:port1,host2:port2,...。客户端使用这些地址来连接 Kafka 集群。
		*/
		"bootstrap.servers": kUrl,

		/*
			group.id
				说明: 这是一个用于标识消费者组的配置项。消费者组是 Kafka 中的一种机制，它允许多个消费者共享同一个主题的消息。
				所有属于同一个消费者组的消费者会共同消费主题中的消息，每条消息只会被组内的一个消费者处理。
				与 Kafka 的配置关系: 这是 Kafka 的标准配置，消费者组的概念是 Kafka 的核心特性之一。所有消费者都需要指定 group.id。
		*/
		"group.id": "0",

		/*
			auto.offset.reset
				earliest: 从最早的消息开始消费。
				latest: 从最新的消息开始消费。
				none: 如果没有找到位移，抛出错误。
		*/
		"auto.offset.reset": "earliest", // 从最早的消息开始消费。
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{config.KafkaTopic, "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		panic(err)
	}

	// 可以使用信号处理程序或类似程序将其设置为 false 以中断循环。
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// 客户端将自动尝试从所有错误中恢复。
			// 超时不被视为错误，因为它是由 ReadMessage 在没有消息的情况下引发的。
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
