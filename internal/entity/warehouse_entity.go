package entity

import (
	"time"

	"github.com/google/uuid"
)

// Warehouse is a struct that represents a warehouse entity
type Warehouse struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"column:name"`
	Status    bool      `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
}

func (u *Warehouse) TableName() string {
	return "warehouses"
}
