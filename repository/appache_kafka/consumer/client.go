package consumer

import "github.com/IBM/sarama"

func NewConsumerClient(addrs []string, groupName string, config *sarama.Config) (sarama.ConsumerGroup, error) {
	return sarama.NewConsumerGroup(addrs, groupName, config)
}
