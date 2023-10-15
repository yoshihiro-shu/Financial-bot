// example of consumer
package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/yoshihiro-shu/financial-bot/internal/logger"
	"github.com/yoshihiro-shu/financial-bot/repository/appache_kafka/consumer"
)

func main() {
	logger := logger.NewSlog()
	brokers := []string{"localhost:9092"}
	group := "notifications"
	topic := "news"

	client, err := consumer.NewConsumerGroupHandler(brokers, group)
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
				log.Fatalf("Error from consumer: %v", err)
			}
		}
	}()

	<-sigchan
	client.Close()
}
