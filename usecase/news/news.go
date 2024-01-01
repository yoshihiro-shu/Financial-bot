package news

import (
	"context"
	"database/sql"
	"time"

	"github.com/yoshihiro-shu/financial-bot/entity"
	"github.com/yoshihiro-shu/financial-bot/internal/finhub"

	repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"
)

func (svc *Service) MarketNews(ctx context.Context, apiKey string) error {
	client := finhub.NewClient(apiKey)
	res, err := client.GetMarketNews()
	if err != nil {
		return err
	}

	categories := make(map[string]int32)
	c, err := svc.Repository.ListCategory(ctx)
	if err != nil {
		return err
	}
	for _, v := range c {
		categories[v.Name] = v.ID
	}

	providers := make(map[string]int32)
	p, err := svc.Repository.ListProvider(ctx)
	if err != nil {
		return err
	}
	for _, v := range p {
		providers[v.Name] = v.ID
	}

	now := time.Now()
	for _, v := range res {
		if _, ok := categories[v.GetCategory()]; !ok {
			category, err := svc.Repository.CreateCategory(ctx, v.GetCategory())
			if err != nil {
				return err
			}
			categories[v.GetCategory()] = category.ID

		}
		if _, ok := providers[v.GetSource()]; !ok {
			provider, err := svc.Repository.CreateProvider(ctx, v.GetSource())
			if err != nil {
				return err
			}
			providers[v.GetSource()] = provider.ID
		}

		// すでに存在している場合はスキップ
		_, err := svc.Repository.GetNewsByTitle(ctx, v.GetHeadline())
		if err != sql.ErrNoRows {
			if err != nil {
				return err
			}
			continue
		}
		_, err = svc.CreateNews(ctx, entity.News{
			Title:       v.GetHeadline(),
			Description: v.GetSummary(),
			Link:        v.GetUrl(),
			Thumbnail:   v.GetImage(),

			PublishedAt: time.Unix(int64(v.GetDatetime()), 0),
			CreatedAt:   now,
			UpdatedAt:   now,
		}, providers[v.GetSource()], categories[v.GetCategory()])
		if err != nil {
			return err
		}
	}
	return nil
}

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
