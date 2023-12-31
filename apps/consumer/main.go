// example of consumer
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/IBM/sarama"
	"github.com/yoshihiro-shu/financial-bot/internal/logger"
	kafka "github.com/yoshihiro-shu/financial-bot/repository/apache_kafka"
)

var (
	brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	group   = os.Getenv("KAFKA_CONSUMER_GROUP")
)

func main() {
	logger := logger.NewSlog()
	group = "notifications"

	client, err := kafka.NewCounsumer(brokers, group)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating consumer group client: %v", err))
	}

	topics, err := client.Topics()
	if err != nil {
		logger.Error(fmt.Sprintf("Error listing topics: %v", err))
	}

	ctx := context.Background()
	handler := kafka.ConsumerGroupHandler{}

	// シグナルを待機（Ctrl+Cで終了）
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	logger.Info("start consuming")

	go func() {
		for {
			err := client.Group().Consume(ctx, topics, handler)
			if err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				logger.Error(fmt.Sprintf("Error from consumer: %v", err))
			}
		}
	}()

	<-sigchan
	logger.Info("Gracefully shutting down")
	client.Close()
}
