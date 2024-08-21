package converter

import (
	"edot-monorepo/services/product-service/internal/entity"
	"edot-monorepo/services/product-service/internal/model"
)

func ProductToResponse(product entity.Product) *model.Product {

	return &model.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

}

func ProductListToProductDetailList(products []entity.Product) []*model.Product {
	productResponse := make([]*model.Product, len(products))

	for i, product := range products {

		productResponse[i] = ProductToResponse(product)
	}

	return productResponse
}
