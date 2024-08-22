package repository

import (
	"edot-monorepo/services/product-service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

func (r *ProductRepository) FindAll(db *gorm.DB, data *[]entity.Product) error {
	return db.Find(data).Error
}
