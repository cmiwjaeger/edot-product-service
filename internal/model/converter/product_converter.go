package converter

import (
	"edot-monorepo/services/product-service/internal/entity"
	"edot-monorepo/services/product-service/internal/model"
)

func ProductToResponse(product entity.Product) *model.ProductResponse {

	return &model.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

}

func ProductListToProductDetailList(products []entity.Product) []*model.ProductResponse {
	productResponse := make([]*model.ProductResponse, len(products))

	for i, product := range products {

		productResponse[i] = ProductToResponse(product)
	}

	return productResponse
}
