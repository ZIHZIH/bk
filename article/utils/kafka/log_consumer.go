package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
	"wzh/article/config"
	"wzh/article/infra"
)

type Consumer struct {
	reader *kafka.Reader
	logger *log.Logger
}

func NewConsumer(logger *log.Logger) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{config.Config.KafkaConfig.Addr},
		GroupID:        "log-group-id",
		Topic:          config.Config.KafkaConfig.Topic,
		Partition:      0,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second 要 加 GroupID 才能自动提交
	})

	return &Consumer{reader: reader, logger: logger}
}

func (c *Consumer) Run() {
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		infra.Logger.Println(string(m.Value))
	}
}
