package controller

import (
	"edot-monorepo/product-service/internal/model"
	"edot-monorepo/product-service/internal/usecase"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	ProductListUseCase *usecase.ProductListUseCase
	Log                *logrus.Logger
	Validate           *validator.Validate
}

func NewProductController(productListUseCase *usecase.ProductListUseCase, log *logrus.Logger, validate *validator.Validate) *ProductController {
	return &ProductController{
		ProductListUseCase: productListUseCase,
		Log:                log,
		Validate:           validate,
	}
}

func (c *ProductController) List(ctx *fiber.Ctx) error {
	request, err := parseQueryToModel(ctx)
	if err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.ProductListUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to list products : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]*model.Product]{
		Data: response,
	})

}

func parseQueryToModel(ctx *fiber.Ctx) (*model.ProductListRequest, error) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		return nil, err
	}

	return &model.ProductListRequest{
		QueryListRequest: model.QueryListRequest{
			Keyword: ctx.Query("keyword"),
			Page:    page,
			Size:    size,
		},
	}, nil
}
