package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/yoshihiro-shu/financial-bot/internal/logger"
	kafka "github.com/yoshihiro-shu/financial-bot/repository/apache_kafka"
)

type server struct {
	consumer *kafka.Consumer
	logger   *slog.Logger
}

func NewServer(conf *sarama.Config) (*server, error) {
	client, err := kafka.NewCounsumer(brokers, group, conf)
	if err != nil {
		return nil, err
	}
	return &server{
		consumer: client,
		logger:   logger.NewSlog(),
	}, nil
}

func (s *server) Run() {
	client, err := kafka.NewCounsumer(brokers, group, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating consumer group client: %v", err))
	}

	s.logger.Info("start consuming")
	ctx := context.Background()
	go func() {
		for {
			err := client.Group().Consume(ctx, []string{topics}, kafka.ConsumerGroupHandler{})
			if err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				s.logger.Error(fmt.Sprintf("Error from consumer: %v", err))
			}
		}
	}()
}

func (s *server) Close() {
	s.logger.Info("Gracefully shutting down...")
	s.consumer.Close()
}
