package model

import "github.com/google/uuid"

type ProductResponse struct {
	ID    uuid.UUID
	Name  string
	Price float64
}

type ProductWithWarehouseResponse struct {
	ProductResponse
	Warehouse WarehouseResponse
}

type ProductListRequest struct {
	QueryListRequest
}
