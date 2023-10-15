package test_container_test

import (
	"context"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/yoshihiro-shu/financial-bot/repository/test_container"
)

func TestAppacheKafkaContainer(t *testing.T) {
	ctx := context.Background()
	container, err := test_container.StartAppacheKafkaContainer(ctx)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	})

	// p, err := producer.NewProducer(container.Addrs)
	// if err != nil {
	// 	t.Errorf("failed to connect producer: %s, err: %s", container.Addrs, err)
	// }

	// t.Run("test send message", func(t *testing.T) {
	// 	err := p.SendMessages("news", []byte("test"), []byte("test"))
	// 	if err != nil {
	// 		t.Errorf("failed to send message: %s", err)
	// 	}
	// })

	// Kafkaプロデューサーの設定
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Version = sarama.V3_5_0_0
	producer, err := sarama.NewSyncProducer(container.Addrs, config)
	if err != nil {
		t.Fatal(err)
	}
	defer producer.Close()

	// メッセージの送信
	topic := "test-topic"
	message := "Hello, Kafka!"
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})
	if err != nil {
		t.Fatal(err)
	}

	// Kafkaコンシューマーの設定
	consumer, err := sarama.NewConsumer(container.Addrs, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer consumer.Close()

	// メッセージの受信
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		t.Fatal(err)
	}
	defer partitionConsumer.Close()

	select {
	case msg := <-partitionConsumer.Messages():
		if string(msg.Value) != message {
			t.Fatalf("expected %s, got %s", message, string(msg.Value))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for the message")
	}
}
