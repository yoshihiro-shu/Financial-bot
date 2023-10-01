package svc

import (
	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
