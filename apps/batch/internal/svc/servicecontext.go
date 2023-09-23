package svc

import (
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
