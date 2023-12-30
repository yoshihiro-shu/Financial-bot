package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Consumer struct {
	client sarama.Client
	group  sarama.ConsumerGroup
}

func NewCounsumer(brokers []string, group string) (*Consumer, error) {
	client, err := sarama.NewClient(brokers, DefaultConfig())
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

func (c *Consumer) Topics() ([]string, error) {
	ts, err := c.client.Topics()
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, fmt.Errorf("no topics found")
	}
	topics := make([]string, len(ts)-1)
	for i, t := range ts {
		if t == "__consumer_offsets" {
			topics = append(ts[:i], ts[i+1:]...)
		}
	}
	return topics, nil
}

type ConsumerGroupHandler struct {
	sarama.ConsumerGroup
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		fmt.Printf("Message key:%s value:%s\n", msg.Key, msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}
