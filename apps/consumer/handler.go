package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

type consumerGroupHandler struct {
	sarama.ConsumerGroup
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		fmt.Printf("Message key:%s value:%s\n", msg.Key, msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}
