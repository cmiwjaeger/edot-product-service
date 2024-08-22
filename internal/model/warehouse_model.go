package model

import "github.com/google/uuid"

type WarehouseResponse struct {
	ID   uuid.UUID
	Name string
}
