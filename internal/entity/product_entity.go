package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ShopID    uuid.UUID `gorm:"column:shop_id"`
	Shop      Shop      `gorm:"foreignKey:ShopID"`
	Name      string    `gorm:"column:name"`
	Price     float64   `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
}

func (u *Product) TableName() string {
	return "products"
}
