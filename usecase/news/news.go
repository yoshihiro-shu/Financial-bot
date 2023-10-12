package news

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/entity"
)

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
