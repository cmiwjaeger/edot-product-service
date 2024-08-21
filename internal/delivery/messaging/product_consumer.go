package messaging

import (
	"edot-monorepo/services/product-service/internal/model"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type ProductConsumer struct {
	Log *logrus.Logger
}

func NewProductConsumer(log *logrus.Logger) *ProductConsumer {
	return &ProductConsumer{
		Log: log,
	}
}

func (c ProductConsumer) Consume(message *kafka.Message) error {
	ContactEvent := new(model.Product)
	if err := json.Unmarshal(message.Value, ContactEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Contact event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic contacts with event: %v from partition %d", ContactEvent, message.TopicPartition.Partition)
	return nil
}
