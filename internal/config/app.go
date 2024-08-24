package config

import (
	"edot-monorepo/services/product-service/internal/delivery/http/controller"
	"edot-monorepo/services/product-service/internal/delivery/http/route"
	"edot-monorepo/services/product-service/internal/gateway/messaging"
	repository "edot-monorepo/services/product-service/internal/repository/gorm"
	"edot-monorepo/services/product-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
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

	Reader *kafka.Reader
	Writer *kafka.Writer
}

func Bootstrap(config *BootstrapConfig) {

	producer := messaging.NewProducer(config.Writer, config.Log)

	productRepository := repository.NewProductRepository(config.Log)

	productUseCase := usecase.NewProductUseCase(config.DB, config.Log, productRepository, config.Validate, producer)

	productCreateUseCase := usecase.NewProductCreateUseCase(productUseCase)
	productListUseCase := usecase.NewProductListUseCase(productUseCase)
	productController := controller.NewProductController(productListUseCase, productCreateUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:               config.App,
		ProductController: productController,
	}

	routeConfig.Setup()
}
