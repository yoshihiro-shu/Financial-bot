package producer

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	sarama.SyncProducer
}

func NewProducer(addrs []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V3_5_0_0

	producer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
		return nil, err
	}

	return &Producer{producer}, nil
}

func (p *Producer) SendMessages(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	_, _, err := p.SyncProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
