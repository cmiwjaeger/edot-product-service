package messaging

import (
	"context"
	"edot-monorepo/services/product-service/internal/entity"
	"edot-monorepo/shared/events"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductConsumer struct {
	Log       *logrus.Logger
	DB        *gorm.DB
	Validator *validator.Validate
}

func NewProductConsumer(log *logrus.Logger, db *gorm.DB, validate *validator.Validate) *ProductConsumer {
	return &ProductConsumer{
		Log:       log,
		DB:        db,
		Validator: validate,
	}
}

func (c ProductConsumer) ConsumeShopCreated(message *kafka.Message) error {
	event := new(events.ShopCreatedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		c.Log.WithError(err).Error("error unmarshalling ShopCreatedEvent event")
		return err
	}
	data := &entity.Shop{
		Name:    event.Name,
		Address: event.Address,
	}

	err := c.DB.Create(data).Error
	if err != nil {
		c.Log.WithError(err).Error("error insert into db")
	}

	c.Log.Infof("Received topic  with event: %v from partition %d", event, message.TopicPartition.Partition)
	return nil
}

func (c ProductConsumer) ConsumeWarehouseCreated(message *kafka.Message, ctx context.Context) error {
	event := new(events.WarehouseCreatedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		c.Log.WithError(err).Error("error unmarshalling WarehouseCreatedEvent")
		return err
	}
	data := &entity.Warehouse{
		Name:   event.Name,
		Status: event.Status,
	}

	err := c.DB.Create(data).Error
	if err != nil {
		c.Log.WithError(err).Error("error insert into db")
	}

	c.Log.Infof("Received topic warehouse with event: %v from partition %d", event, message.TopicPartition.Partition)
	return nil
}
