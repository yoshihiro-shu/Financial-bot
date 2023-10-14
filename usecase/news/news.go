package news

import (
	"context"
	"database/sql"

	"github.com/yoshihiro-shu/financial-bot/entity"
	repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"
)

func (svc *Service) CreateNews(ctx context.Context, news entity.News, providerId int32, categoryId int32) (*entity.News, error) {
	res, err := svc.Repository.CreateNews(ctx, repository.CreateNewsParams{
		Title:       news.Title,
		Description: sql.NullString{String: news.Description, Valid: false},
		Link:        news.Link,
		Thumbnail:   sql.NullString{String: news.Thumbnail, Valid: false},
		Score:       news.Score,
		PublishedAt: news.PublishedAt,
		ProviderID:  sql.NullInt32{Int32: providerId, Valid: true},
		CategoryID:  sql.NullInt32{Int32: categoryId, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &entity.News{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description.String,
		Link:        res.Link,
		Thumbnail:   res.Thumbnail.String,
		Score:       res.Score,
		PublishedAt: res.PublishedAt,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}

func (svc *Service) DeleteNews(ctx context.Context, id int32) error {
	return svc.Repository.DeleteNews(ctx, id)
}

func (svc *Service) GetNews(ctx context.Context, id int32) (*entity.News, error) {
	news, err := svc.Repository.GetNews(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.News{
		ID:          news.ID,
		Title:       news.Title,
		Description: news.Description.String,
		Link:        news.Link,
		Thumbnail:   news.Thumbnail.String,
		Score:       news.Score,
		PublishedAt: news.PublishedAt,
		CreatedAt:   news.CreatedAt,
		UpdatedAt:   news.UpdatedAt,
	}, nil
}

func (svc *Service) ListNews(ctx context.Context) ([]entity.News, error) {
	res, err := svc.Repository.ListNews(ctx)
	if err != nil {
		return nil, err
	}

	news := make([]entity.News, 0)
	for _, v := range res {
		news = append(news, entity.News{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description.String,
			Link:        v.Link,
			Thumbnail:   v.Thumbnail.String,
			Score:       v.Score,
			PublishedAt: v.PublishedAt,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return news, nil
}
