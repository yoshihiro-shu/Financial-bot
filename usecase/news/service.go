package news

import repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"

type Service struct {
	Repository *repository.Queries
}

func NewService(repo *repository.Queries) *Service {
	return &Service{
		Repository: repo,
	}
}
