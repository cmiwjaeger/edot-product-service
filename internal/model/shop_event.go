package model

import "github.com/google/uuid"

type ShopEvent struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"string"`
}
