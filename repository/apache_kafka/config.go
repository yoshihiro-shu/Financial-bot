package kafka

import "github.com/IBM/sarama"

func consumerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_5_0_0 // Kafkaのバージョンに合わせて変更
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	return config
}

func producerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V3_5_0_0
	return config
}
