package model

import "github.com/google/uuid"

type ProductCreatedEvent struct {
	ID     uuid.UUID `json:"uuid"`
	Name   string    `json:"name"`
	Stock  int64     `json:"stock"`
	ShopID uuid.UUID `json:"shop_id"`
	Shop   ShopEvent `json:"shop"`
}
