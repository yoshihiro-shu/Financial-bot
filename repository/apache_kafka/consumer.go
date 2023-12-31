package kafka

import (
	"github.com/IBM/sarama"
)

type Consumer struct {
	client sarama.Client
	group  sarama.ConsumerGroup
}

func NewCounsumer(brokers, group string, conf *sarama.Config) (*Consumer, error) {
	if conf == nil {
		conf = consumerConfig()
	}
	client, err := sarama.NewClient([]string{brokers}, conf)
	if err != nil {
		return nil, err
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(group, client)
	if err != nil {
		return nil, err
	}
	return &Consumer{client, consumerGroup}, nil
}

func (c *Consumer) Group() sarama.ConsumerGroup {
	return c.group
}

func (c *Consumer) Close() error {
	return c.client.Close()
}
