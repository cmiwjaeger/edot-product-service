package messaging

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type ConsumerHandler func(ctx context.Context, message *kafka.Message) error

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(reader *kafka.Reader) *Consumer {
	return &Consumer{
		Reader: reader,
	}
}

func (c *Consumer) Consume(ctx context.Context, handlers map[string]ConsumerHandler, log *logrus.Logger) (err error) {
	for {
		m, errM := c.Reader.ReadMessage(ctx)
		if errM != nil {
			log.Printf("Failed to read message: %v\n", err)
			return
		}
		log.Printf("Received message: %s = %s\n", string(m.Key), string(m.Value))

		if handler, exist := handlers[m.Topic]; exist {
			err = handler(ctx, &m)
			if err != nil {
				log.Errorf("Failed to process message on topic %s: %v", m.Topic, err)
			} else {
				err = c.Reader.CommitMessages(ctx, m)
				if err != nil {
					log.Fatalf("Failed to commit message on topic %s: %v", m.Topic, err)
				}
			}
		} else {
			log.Warnf("No handler found for topic %s", m.Topic)
		}
	}
}
