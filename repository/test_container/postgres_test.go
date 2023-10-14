package test_container_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/financial-bot/repository/test_container"
)

func TestStartPostgresContainer(t *testing.T) {
	ctx := context.Background()
	container, err := test_container.StartPostgresContainer(ctx)
	if err != nil {
		t.Errorf("error is %s", err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	})

	err = container.InitMigration(ctx)
	if err != nil {
		t.Errorf("error is %s", err)
	}

	err = container.InitTestData(ctx)
	if err != nil {
		t.Errorf("error is %s", err)
	}

	db, err := sql.Open("postgres", container.URI)
	if err != nil {
		t.Errorf("error is %s", err)
	}
	defer db.Close()

	// send ping
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
