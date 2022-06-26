package messaging

import (
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucasd-coder/classroom/internal/pkg/kafka"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
)

type Consumer struct {
	handler func(m *ckafka.Message) error
	config  *ckafka.ConfigMap
}

func New(handler func(m *ckafka.Message) error, config *ckafka.ConfigMap) *Consumer {
	return &Consumer{
		handler: handler,
		config:  config,
	}
}

func (c *Consumer) Start() {
	time.Sleep(time.Second * 2)

	msgChan := make(chan *ckafka.Message)

	topics := []string{os.Getenv("KAFKA_PURCHASES_TOPIC")}

	consumer := kafka.NewConsumer(c.config, topics)

	go consumer.Consume(msgChan)

	logger.Log.Info("Kafka consumer has been started")

	for msg := range msgChan {

		err := c.handler(msg)
		if err != nil {
			logger.Log.Error(err)
		}
	}
}
