package usecase

import (
	"context"
	"edot-monorepo/services/product-service/internal/entity"
	"edot-monorepo/services/product-service/internal/gateway/messaging"
	"edot-monorepo/services/product-service/internal/model"
	"edot-monorepo/services/product-service/internal/model/converter"

	"github.com/gofiber/fiber/v2"
)

type ProductCreateUseCase struct {
	*ProductBaseUseCase
	ProductCreatedProducer *messaging.ProductProducer[model.Event]
}

func NewProductCreateUseCase(productBaseUseCase *ProductBaseUseCase, productCreatedProducer *messaging.ProductProducer[model.Event]) *ProductCreateUseCase {
	return &ProductCreateUseCase{
		ProductBaseUseCase:     productBaseUseCase,
		ProductCreatedProducer: productCreatedProducer,
	}
}

func (c *ProductCreateUseCase) Exec(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	product := &entity.Product{
		Name: request.Name,
	}

	if err := c.ProductRepository.Create(tx, product); err != nil {
		c.Log.Warnf("Failed create product to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	event := converter.ProductToEvent(product)
	if err := c.ProductCreatedProducer.SendAsync(event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}
