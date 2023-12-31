// example of consumer
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/yoshihiro-shu/financial-bot/internal/logger"
	kafka "github.com/yoshihiro-shu/financial-bot/repository/apache_kafka"
)

// Sarama configuration options
var (
	brokers  = ""
	version  = ""
	group    = ""
	topics   = ""
	assignor = ""
	oldest   = true
	verbose  = false
)

func init() {
	flag.StringVar(&brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", "", "Kafka consumer group definition")
	flag.StringVar(&version, "version", sarama.DefaultVersion.String(), "Kafka cluster version")
	flag.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func main() {
	s, err := newServer()
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating server: %v", err))
		return
	}

	s.Run()

	// シグナルを待機（Ctrl+Cで終了）
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	<-sigchan
	s.Close()
}

type server struct {
	consumer *kafka.Consumer
	logger   *slog.Logger
}

func newServer() (*server, error) {
	client, err := kafka.NewCounsumer(brokers, group, nil)
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
