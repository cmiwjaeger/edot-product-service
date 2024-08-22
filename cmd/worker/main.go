package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"edot-monorepo/services/product-service/internal/config"
	"edot-monorepo/services/product-service/internal/delivery/messaging"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	// Initialize configurations, logger, and other dependencies
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logger)
	validate := config.NewValidator(viperConfig)

	// Start the service
	logger.Info("Starting worker service")

	// Set up context with cancel for graceful shutdown using signal.NotifyContext
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop() // Ensure stop is called on exit

	// Use a WaitGroup to wait for the consumer to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Run the product consumer in a separate goroutine
	go func() {
		defer wg.Done()
		RunProductConsumer(logger, db, validate, viperConfig, ctx)
	}()

	// Wait for context cancellation (signal received)
	<-ctx.Done()
	logger.Info("Received shutdown signal, waiting for goroutines to finish")

	// Wait for the consumer to finish processing
	wg.Wait()
	logger.Info("Worker service has shut down gracefully")
}

func RunProductConsumer(logger *logrus.Logger, db *gorm.DB, validate *validator.Validate, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup product consumer")
	consumer := config.NewKafkaConsumer(viperConfig, logger)
	handler := messaging.NewProductConsumer(logger, db, validate)

	messaging.ConsumeTopic(ctx, consumer, "shop_created", logger, handler.ConsumeShopCreated)
	messaging.ConsumeTopic(ctx, consumer, "warehouse_created", logger, handler.ConsumeShopCreated)

}
