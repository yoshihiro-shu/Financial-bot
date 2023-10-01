package logic

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotificationFinancialNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotificationFinancialNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotificationFinancialNewsLogic {
	return &NotificationFinancialNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotificationFinancialNewsLogic) NotificationFinancialNews(req *types.NotificationFinancialNewsReq) (resp *types.NotificationFinancialNewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
