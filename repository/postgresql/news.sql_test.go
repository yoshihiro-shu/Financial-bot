package postgresql_test

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yoshihiro-shu/financial-bot/repository/postgresql"
	"github.com/yoshihiro-shu/financial-bot/repository/test_container"
)

func TestGetNews(t *testing.T) {
	ctx := context.Background()
	container := test_container.Postgres
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	})

	db, err := sql.Open("postgres", container.URI)
	if err != nil {
		t.Errorf("error is %s", err)
	}
	defer db.Close()

	tests := []struct {
		Name string
		Args int32
		Want postgresql.News
	}{
		{
			Name: "Get News ID 1",
			Args: 1,
			Want: postgresql.News{
				ID:          1,
				Title:       "News Title 1",
				Description: sql.NullString{String: "News Description 1", Valid: true},
				Link:        "News Link 1",
				Thumbnail:   sql.NullString{String: "News Thumbnail 1", Valid: true},
				Score:       100,
				PublishedAt: parseTime("2023-10-13 12:00:00 +0900"),
				ProviderID:  sql.NullInt32{Int32: 1, Valid: true},
				CategoryID:  sql.NullInt32{Int32: 1, Valid: true},
			},
		},
		{
			Name: "Get News ID 2",
			Args: 2,
			Want: postgresql.News{
				ID:          2,
				Title:       "News Title 2",
				Description: sql.NullString{String: "News Description 2", Valid: true},
				Link:        "News Link 2",
				Thumbnail:   sql.NullString{String: "News Thumbnail 2", Valid: true},
				Score:       90,
				PublishedAt: parseTime("2023-10-13 12:00:00 +0900"),
				ProviderID:  sql.NullInt32{Int32: 2, Valid: true},
				CategoryID:  sql.NullInt32{Int32: 2, Valid: true},
			},
		},
		{
			Name: "Get News ID 3",
			Args: 3,
			Want: postgresql.News{
				ID:          3,
				Title:       "News Title 3",
				Description: sql.NullString{String: "News Description 3", Valid: true},
				Link:        "News Link 3",
				Thumbnail:   sql.NullString{String: "News Thumbnail 3", Valid: true},
				Score:       80,
				PublishedAt: parseTime("2023-10-13 12:00:00 +0900"),
				ProviderID:  sql.NullInt32{Int32: 3, Valid: true},
				CategoryID:  sql.NullInt32{Int32: 3, Valid: true},
			},
		},
	}

	repo := postgresql.New(db)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := repo.GetNews(ctx, test.Args)
			if err != nil {
				t.Errorf("error is %s", err)
			}
			assert.Equal(t, test.Want.Title, res.Title)
			assert.Equal(t, test.Want.Description, res.Description)
			assert.Equal(t, test.Want.Link, res.Link)
			assert.Equal(t, test.Want.Thumbnail, res.Thumbnail)
			assert.Equal(t, test.Want.Score, res.Score)
			assert.Equal(t, test.Want.PublishedAt, res.PublishedAt)
			assert.Equal(t, test.Want.ProviderID, res.ProviderID)
			assert.Equal(t, test.Want.CategoryID, res.CategoryID)
		})
	}

}

func parseTime(unixTime string) time.Time {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("Failed to load location: %s", err)
	}
	const layout = "2006-01-02 15:04:05 -0700"
	res, err := time.Parse(layout, unixTime)
	if err != nil {
		log.Fatalf("Failed to parse time: %s", err)
	}

	return res.In(location)
}
