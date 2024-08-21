package usecase

import (
	"context"
	"edot-monorepo/product-service/internal/entity"
	"edot-monorepo/product-service/internal/model"
	"edot-monorepo/product-service/internal/model/converter"
	repository "edot-monorepo/product-service/internal/repository/gorm"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductListUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	ProductRepository *repository.ProductRepository
}

func NewProductListUseCase(db *gorm.DB, log *logrus.Logger, repo *repository.ProductRepository) *ProductListUseCase {
	return &ProductListUseCase{
		DB:                db,
		ProductRepository: repo,
	}
}

func (u *ProductListUseCase) Exec(ctx context.Context, request *model.ProductListRequest) ([]*model.Product, error) {
	products := make([]entity.Product, 0)

	err := u.ProductRepository.FindAll(u.DB, &products)
	if err != nil {
		log.Errorf("%v", err)
	}

	return converter.ProductListToProductDetailList(products), nil
}
