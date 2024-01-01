package news_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"
	"github.com/yoshihiro-shu/financial-bot/usecase/news"
)

var apiKey = os.Getenv("FINHUB_API_KEY")

// TODO REFACTOR
func TestMarketNews(t *testing.T) {
	if apiKey == "" {
		t.Skip()
	}
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=test sslmode=disable")
	assert.Nil(t, err)
	defer db.Close()
	if err != nil {
		t.Errorf("error is %s", err)
	}
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
