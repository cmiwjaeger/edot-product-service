package usecase

import (
	"context"
	"edot-monorepo/services/product-service/internal/entity"
	"edot-monorepo/services/product-service/internal/model"
	"edot-monorepo/services/product-service/internal/model/converter"

	"github.com/gofiber/fiber/v2/log"
)

type ProductListUseCase struct {
	*ProductBaseUseCase
}

func NewProductListUseCase(productBaseUseCase *ProductBaseUseCase) *ProductListUseCase {

	return &ProductListUseCase{
		productBaseUseCase,
	}
}

func (u *ProductListUseCase) Exec(ctx context.Context, request *model.ProductListRequest) ([]*model.ProductResponse, error) {
	products := make([]entity.Product, 0)

	err := u.ProductRepository.FindAll(u.DB, &products)
	if err != nil {
		log.Errorf("%v", err)
	}

	return converter.ProductListToProductDetailList(products), nil
}
