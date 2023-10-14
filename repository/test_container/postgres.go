package test_container

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/yoshihiro-shu/financial-bot/internal/testutils"
)

const migrationsDir = "/db/migrations"

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

func (c *postgresDBContainer) InitMigration(ctx context.Context) error {
	db, err := sql.Open("postgres", c.URI)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	rootDir, _ := testutils.GetProjectRoot()
	if err := goose.Up(db, rootDir+migrationsDir); err != nil {
		return err
	}
	return nil
}
