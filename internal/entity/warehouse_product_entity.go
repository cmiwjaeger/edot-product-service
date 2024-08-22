package entity

import "github.com/google/uuid"

type WarehouseProduct struct {
	ID          uuid.UUID `gorm:"column:id"`
	ProductID   uuid.UUID `gorm:"column:product_id"`
	Product     Product   `gorm:"references:ProductID"`
	WarehouseID uuid.UUID `gorm:"column:warehouse_id"`
	Warehouse   Warehouse `gorm:"references:WarehouseID"`
}

func (u *WarehouseProduct) TableName() string {
	return "warehouse_products"
}
