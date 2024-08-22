package converter

import (
	"edot-monorepo/services/product-service/internal/entity"
	"edot-monorepo/services/product-service/internal/model"
	"edot-monorepo/shared/events"
)

func ProductToEvent(item *entity.Product) *events.ProductCreatedEvent {
	return &events.ProductCreatedEvent{
		ID:   item.ID,
		Name: item.Name,
	}
}

func ProductToResponse(product *entity.Product) *model.ProductResponse {

	return &model.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func ProductListToProductDetailList(products []entity.Product) []*model.ProductResponse {
	productResponse := make([]*model.ProductResponse, len(products))

	for i, product := range products {

		productResponse[i] = ProductToResponse(&product)
	}

	return productResponse
}
