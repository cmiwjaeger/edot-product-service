package messaging

import (
	"edot-monorepo/services/product-service/internal/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type ProductProducer[T model.Event] struct {
	Producer[T]
}

func NewProductProducer[T model.Event](topic string, producer *kafka.Producer, log *logrus.Logger) *ProductProducer[model.Event] {

	return &ProductProducer[model.Event]{
		Producer: Producer[model.Event]{
			Producer: producer,
			Topic:    topic,
			Log:      log,
		},
	}
}
