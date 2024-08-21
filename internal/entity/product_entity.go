package entity

import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID `gorm:"column:id;primaryKey"`
	Name  string    `gorm:"column:name"`
	Price float64   `gorm:"column:price"`
}

func (u *Product) TableName() string {
	return "products"
}
