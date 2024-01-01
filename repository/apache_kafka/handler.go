package kafka

import (
	"fmt"
	"net/http"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct {
	sarama.ConsumerGroup
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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

func handleRequest(msg *sarama.ConsumerMessage) (*http.Response, error) {
	switch string(msg.Key) {
	case http.MethodGet:
		return http.DefaultClient.Get(string(string(msg.Value)))
	}
	return nil, nil
}

func handleDefault(msg *sarama.ConsumerMessage) {
	fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
	fmt.Printf("Message key:%s value:%s\n", msg.Key, msg.Value)
}
