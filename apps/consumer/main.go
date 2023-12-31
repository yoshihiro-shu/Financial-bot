package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	brokers  = ""
	version  = ""
	group    = ""
	topics   = ""
	assignor = ""
	oldest   = true
	verbose  = false
)

func init() {
	flag.StringVar(&brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", "", "Kafka consumer group definition")
	flag.StringVar(&version, "version", sarama.DefaultVersion.String(), "Kafka cluster version")
	flag.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func main() {
	conf, err := ParseConfig()
	if err != nil {
		panic(err)
	}

	s, err := NewServer(conf)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating server: %v", err))
		return
	}

	s.Run()

	// シグナルを待機（Ctrl+Cで終了）
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	<-sigchan
	s.Close()
}
