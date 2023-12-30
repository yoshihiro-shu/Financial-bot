package consumer

import "github.com/IBM/sarama"

func DefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_5_0_0 // Kafkaのバージョンに合わせて変更
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	return config
}
