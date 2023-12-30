// example of consumer
package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/IBM/sarama"
	"github.com/yoshihiro-shu/financial-bot/internal/logger"
	"github.com/yoshihiro-shu/financial-bot/repository/appache_kafka/consumer"
)

var (
	brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

func main() {
	logger := logger.NewSlog()
	group := "notifications"
	topic := "news"

	config := consumer.DefaultConfig()
	client, err := consumer.NewConsumerClient(brokers, group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}
	ctx := context.Background()
	handler := consumer.ConsumerGroupHandler{}

	// シグナルを待機（Ctrl+Cで終了）
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	logger.Info("start consuming")

	// Consumer Groupでのメッセージの消費を開始
	go func() {
		for {
			err := client.Consume(ctx, []string{topic}, handler)
			if err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Fatalf("Error from consumer: %v", err)
			}
		}
	}()

	<-sigchan
	client.Close()
}
