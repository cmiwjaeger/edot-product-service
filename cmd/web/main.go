package main

import (
	"context"
	"edot-monorepo/services/product-service/internal/config"
	"fmt"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	kafkaWriter := config.NewKafkaWriter(viperConfig, log)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		Writer:   kafkaWriter,
	})

	// Start Fiber in a goroutine
	go func() {
		webPort := viperConfig.GetInt("web.port")
		if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Channel to listen for interrupt or terminate signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-quit

	// Graceful shutdown
	log.Println("Shutting down gracefully...")

	// Gracefully shutdown Fiber
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

}
