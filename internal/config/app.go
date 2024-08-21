package config

import (
	"edot-monorepo/product-service/internal/delivery/http/controller"
	"edot-monorepo/product-service/internal/delivery/http/route"
	repository "edot-monorepo/product-service/internal/repository/gorm"
	"edot-monorepo/product-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	productRepository := repository.NewProductRepository(config.Log)
	productListUseCase := usecase.NewProductListUseCase(config.DB, config.Log, productRepository)
	productController := controller.NewProductController(productListUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:               config.App,
		ProductController: productController,
	}

	routeConfig.Setup()
}
