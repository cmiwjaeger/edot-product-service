package usecase

import (
	"edot-monorepo/services/product-service/internal/gateway/messaging"
	repository "edot-monorepo/services/product-service/internal/repository/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductBaseUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	ProductRepository *repository.ProductRepository
	Validate          *validator.Validate
	Producer          *messaging.Producer
}

func NewProductUseCase(db *gorm.DB, log *logrus.Logger, repo *repository.ProductRepository, validate *validator.Validate, producer *messaging.Producer) *ProductBaseUseCase {
	return &ProductBaseUseCase{
		DB:                db,
		Log:               log,
		ProductRepository: repo,
		Validate:          validate,
		Producer:          producer,
	}
}
