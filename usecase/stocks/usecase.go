package stocks

import (
	"context"
)

type UseCase interface {
	CreateStocks(ctx context.Context) error
}
