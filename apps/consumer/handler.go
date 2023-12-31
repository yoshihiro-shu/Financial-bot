package main

import (
	"fmt"
	"net/http"

	"github.com/IBM/sarama"
)

type consumerGroupHandler struct {
	sarama.ConsumerGroup
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		switch msg.Topic {
		case "request":
			handleRequest(msg)
		default:
			handleDefault(msg)
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

func handleRequest(msg *sarama.ConsumerMessage) {
	switch string(msg.Key) {
	case http.MethodGet:
		http.Get(string(msg.Value))
	}
}

func handleDefault(msg *sarama.ConsumerMessage) {
	fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
	fmt.Printf("Message key:%s value:%s\n", msg.Key, msg.Value)
}
