package news

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/entity"
)

type UseCase interface {
	GetNews(ctx context.Context, id int32) (*entity.News, error)
}
