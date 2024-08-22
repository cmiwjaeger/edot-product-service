package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"edot-monorepo/services/product-service/internal/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")

	ctx, cancel := context.WithCancel(context.Background())

	go RunProductConsumer(logger, viperConfig, ctx)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}

func RunProductConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup product consumer")
	// productConsumer := config.NewKafkaConsumer(viperConfig, logger)
	// addressHandler := messaging.NewProductConsumer(logger)
	// messaging.ConsumeTopic(ctx, productConsumer, "addresses", logger, addressHandler.Consume)
}
