package news_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"
	"github.com/yoshihiro-shu/financial-bot/repository/test_container"
	"github.com/yoshihiro-shu/financial-bot/usecase/news"
)

var apiKey = os.Getenv("FINHUB_API_KEY")

// TODO REFACTOR
func TestMarketNews(t *testing.T) {
	if apiKey == "" {
		t.Skip()
	}
	ctx := context.Background()
	container, err := test_container.NewPostgresDBContainer(ctx)
	assert.Nil(t, err)
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	})

	db, err := sql.Open("postgres", container.URI)
	assert.Nil(t, err)
	defer db.Close()

	repo := repository.New(db)
	svc := &news.Service{Repository: repo}
	err = svc.MarketNews(context.Background(), apiKey)
	assert.Nil(t, err)
	list, err := repo.ListNews(context.Background())
	assert.Nil(t, err)
	if len(list) == 0 {
		t.Errorf("error is %s", err)
	}
}
