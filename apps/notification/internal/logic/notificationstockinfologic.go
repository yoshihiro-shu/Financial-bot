package logic

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotificationStockInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotificationStockInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotificationStockInfoLogic {
	return &NotificationStockInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotificationStockInfoLogic) NotificationStockInfo(req *types.NotificationStockInfoReq) (resp *types.NotificationStockInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
