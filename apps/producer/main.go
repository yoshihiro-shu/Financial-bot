// example of producer
package main

import (
	"log"

	"github.com/yoshihiro-shu/financial-bot/repository/appache_kafka/producer"
)

func main() {
	producer, err := producer.NewProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	log.Printf("==> start producing %#v\n", producer)
	log.Printf("==> status of producing %#v", producer.TxnStatus())

	topic := "news"
	key := []byte("key")
	value := []byte("value")
	err = producer.SendMessages(topic, key, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success")
}
