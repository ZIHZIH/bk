package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"wzh/article/config"
)

type Producer struct {
	w *kafka.Writer
}

func (p *Producer) Write(ctx context.Context, value string) {
	if err := p.w.WriteMessages(ctx, kafka.Message{Key: []byte("wzh666"), Value: []byte(value)}); err != nil {
		panic(err)
	}
}

func NewProducer() *Producer {
	fmt.Println(config.Config.KafkaConfig)
	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Config.KafkaConfig.Addr),
		Topic:                  config.Config.KafkaConfig.Topic,
		Balancer:               &kafka.LeastBytes{}, // 指定分区的balancer模式为最小字节分布
		RequiredAcks:           kafka.RequireAll,    // ack模式
		AllowAutoTopicCreation: true,
	}
	return &Producer{w: w}
}
