package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	sarama.SyncProducer
}

func NewProducer(addrs []string) (*Producer, error) {
	client, err := sarama.NewClient(addrs, producerConfig())
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		return nil, err
	}
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
		return nil, err
	}

	return &Producer{producer}, nil
}

func (p *Producer) SendMessage(topic string, key, value []byte) error {
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
