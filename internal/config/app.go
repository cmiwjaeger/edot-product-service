package config

import (
	"edot-monorepo/services/product-service/internal/delivery/http/controller"
	"edot-monorepo/services/product-service/internal/delivery/http/route"
	"edot-monorepo/services/product-service/internal/gateway/messaging"
	repository "edot-monorepo/services/product-service/internal/repository/gorm"
	"edot-monorepo/services/product-service/internal/usecase"
	"edot-monorepo/shared/events"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
	Producer *kafka.Producer
}

func Bootstrap(config *BootstrapConfig) {

	productCreatedProducer := messaging.NewProductProducer[*events.WarehouseCreatedEvent]("product_created", config.Producer, config.Log)

	productRepository := repository.NewProductRepository(config.Log)

	productUseCase := usecase.NewProductUseCase(config.DB, config.Log, productRepository, config.Validate)

	productCreateUseCase := usecase.NewProductCreateUseCase(productUseCase, productCreatedProducer)
	productListUseCase := usecase.NewProductListUseCase(productUseCase)
	productController := controller.NewProductController(productListUseCase, productCreateUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:               config.App,
		ProductController: productController,
	}

	routeConfig.Setup()
}
