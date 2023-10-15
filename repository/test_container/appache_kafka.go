package test_container

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/kafka"
)

type AppacheKafkaContainer struct {
	kafka.KafkaContainer
	Addrs []string
}

func StartAppacheKafkaContainer(ctx context.Context) (*AppacheKafkaContainer, error) {
	kafkaContainer, err := kafka.RunContainer(ctx,
		kafka.WithClusterID("test-cluster"),
		testcontainers.WithImage("confluentinc/confluent-local:7.5.1"),
	)
	if err != nil {
		return nil, err
	}

	brokers, err := kafkaContainer.Brokers(ctx)
	if err != nil {
		return nil, err
	}

	return &AppacheKafkaContainer{
		KafkaContainer: *kafkaContainer,
		Addrs:          brokers,
	}, nil
}

// type AppacheKafkaContainer struct {
// 	testcontainers.Container
// 	Addrs []string
// }

// func StartAppacheKafkaContainer(ctx context.Context) (*AppacheKafkaContainer, error) {
// 	req := testcontainers.ContainerRequest{
// 		Image:        "docker.io/bitnami/kafka:3.5",
// 		ExposedPorts: []string{"9092/tcp", "9093/tcp"},
// 		WaitingFor:   wait.ForLog("==> ** Starting Kafka **"),
// 		Env: map[string]string{
// 			// KRaft settings
// 			"KAFKA_CFG_NODE_ID":                  "0",
// 			"KAFKA_CFG_PROCESS_ROLES":            "controller,broker",
// 			"KAFKA_CFG_CONTROLLER_QUORUM_VOTERS": "0@kafka:9093",
// 			// Listeners
// 			"KAFKA_CFG_LISTENERS":                      "PLAINTEXT://:9092,CONTROLLER://:9093",
// 			"KAFKA_CFG_ADVERTISED_LISTENERS":           "PLAINTEXT://localhost:9092",
// 			"KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP": "CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT",
// 			"KAFKA_CFG_CONTROLLER_LISTENER_NAMES":      "CONTROLLER",
// 			"KAFKA_CFG_INTER_BROKER_LISTENER_NAME":     "PLAINTEXT",
// 		},
// 	}

// 	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: req,
// 		Started:          true,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	brokerMappedPort, err := container.MappedPort(ctx, "9092")
// 	if err != nil {
// 		return nil, err
// 	}

// 	plaintextMappedPort, err := container.MappedPort(ctx, "9093")
// 	if err != nil {
// 		return nil, err
// 	}

// 	hostIP, err := container.Host(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	iP, err := container.ContainerIP(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &AppacheKafkaContainer{
// 		Container: container,
// 		Addrs: []string{
// 			fmt.Sprintf("%s:%s", hostIP, brokerMappedPort.Port()),
// 			fmt.Sprintf("%s:%s", iP, brokerMappedPort.Port()),
// 			fmt.Sprintf("%s:%s", hostIP, plaintextMappedPort.Port()),
// 			fmt.Sprintf("%s:%s", iP, plaintextMappedPort.Port()),
// 		},
// 	}, nil
// }
