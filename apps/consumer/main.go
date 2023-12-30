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
	kafka "github.com/yoshihiro-shu/financial-bot/repository/appache_kafka"
)

var (
	brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	group   = os.Getenv("KAFKA_CONSUMER_GROUP")
)

func main() {
	logger := logger.NewSlog()
	group = "notifications"

	config := kafka.DefaultConfig()
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating consumer group client: %v", err))
	}

	consumeGroup, err := sarama.NewConsumerGroupFromClient(group, client)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating consumer group: %v", err))
	}

	ts, err := client.Topics()
	if err != nil {
		logger.Error(fmt.Sprintf("Error listing topics: %v", err))
	}

	topics := make([]string, len(ts)-1)
	for i, t := range ts {
		if t != "__consumer_offsets" {
			topics[i] = t
		}
	}

	ctx := context.Background()
	handler := kafka.ConsumerGroupHandler{}

	// シグナルを待機（Ctrl+Cで終了）
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	logger.Info("start consuming")

	go func() {
		for {
			err := consumeGroup.Consume(ctx, topics, handler)
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
