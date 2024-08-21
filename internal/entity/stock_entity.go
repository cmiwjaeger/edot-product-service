package entity

import (
	"github.com/google/uuid"
)

type Stock struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	WarehouseID uuid.UUID `gorm:"type:uuid;column:warehouse_id"`
	ProductID   uuid.UUID `gorm:"type:uuid:column:product_id"`
	QTY         int64     `gorm:"column:qty"`
}

func (u *Stock) TableName() string {
	return "stock"
}
