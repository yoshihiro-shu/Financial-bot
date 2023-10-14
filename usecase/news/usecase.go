package news

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/entity"
)

type UseCase interface {
	CreateNews(ctx context.Context, news entity.News, providerId int32, categoryId int32) (*entity.News, error)
	DeleteNews(ctx context.Context, id int32) error
	GetNews(ctx context.Context, id int32) (*entity.News, error)
	ListNews(ctx context.Context) ([]entity.News, error)
}
