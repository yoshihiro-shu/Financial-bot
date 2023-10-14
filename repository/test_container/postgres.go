package test_container

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/yoshihiro-shu/financial-bot/internal/testutils"
)

const migrationsDir = "/db/migrations"
const testDataSQLPath = "/db/test/init.sql"

var Postgres *postgresDBContainer

func init() {
	ctx := context.Background()

	psql, err := NewPostgresDBContainer(ctx)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}
	Postgres = psql
}

type postgresDBContainer struct {
	testcontainers.Container
	URI string
}

func NewPostgresDBContainer(ctx context.Context) (*postgresDBContainer, error) {
	c, err := StartPostgresContainer(ctx)
	if err != nil {
		return nil, err
	}
	c.InitMigration(ctx)
	c.InitTestData(ctx)
	return c, nil
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

func (c *postgresDBContainer) InitTestData(ctx context.Context) error {
	db, err := sql.Open("postgres", c.URI)
	if err != nil {
		return err
	}

	rootDir, _ := testutils.GetProjectRoot()

	// SQLファイルを読み込む
	file, err := os.ReadFile(rootDir + testDataSQLPath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %v", err)

	}

	// SQLファイルを実行する
	if _, err := db.Exec(string(file)); err != nil {
		log.Fatalf("Failed to execute SQL: %v", err)
		return fmt.Errorf("failed to execute SQL: %v", err)
	}
	return nil
}
