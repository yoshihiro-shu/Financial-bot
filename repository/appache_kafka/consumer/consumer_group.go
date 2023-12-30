package consumer

import (
	"fmt"

	"github.com/IBM/sarama"
)

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

func NewConsumerGroupHandler(addrs []string, groupName string) (*ConsumerGroupHandler, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_5_0_0 // Kafkaのバージョンに合わせて変更
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	cg, err := sarama.NewConsumerGroup(addrs, groupName, config)
	if err != nil {
		return nil, err
	}
	return &ConsumerGroupHandler{
		ConsumerGroup: cg,
	}, nil
}
