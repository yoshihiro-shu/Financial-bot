package test_container

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type postgresDBContainer struct {
	testcontainers.Container
	URI string
}

func StartPostgresContainer(ctx context.Context) (*postgresDBContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:16.0",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "postgres",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "password",
			"PGUSER":            "postgres",
			"TZ":                "Asia/Tokyo",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("postgresql://postgres:password@%s:%s/postgres?sslmode=disable", hostIP, mappedPort.Port())

	return &postgresDBContainer{Container: container, URI: uri}, nil
}
