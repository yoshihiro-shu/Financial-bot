package logic

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotificationCryptoCurrencyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotificationCryptoCurrencyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotificationCryptoCurrencyInfoLogic {
	return &NotificationCryptoCurrencyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotificationCryptoCurrencyInfoLogic) NotificationCryptoCurrencyInfo(req *types.NotificationCryptoCurrencyInfoReq) (resp *types.NotificationCryptoCurrencyInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
