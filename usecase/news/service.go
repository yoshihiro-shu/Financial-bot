package news

import (
	"log/slog"

	repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"
)

type Service struct {
	Repository *repository.Queries
	logger     *slog.Logger
}

func NewService(repo *repository.Queries, logger *slog.Logger) UseCase {
	return &Service{
		Repository: repo,
		logger:     logger,
	}
}
