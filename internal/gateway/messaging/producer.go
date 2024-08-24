package messaging

import (
	"context"
	"edot-monorepo/shared/events"

	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	Writer *kafka.Writer
	Log    *logrus.Logger
}

func NewProducer(writer *kafka.Writer, log *logrus.Logger) *Producer {
	return &Producer{
		Writer: writer,
		Log:    log,
	}
}

func (p *Producer) Produce(ctx context.Context, topic string, event events.Event) (err error) {

	value, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}
	err = p.Writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   []byte(event.GetId()),
		Value: value,
	})
	return
}
