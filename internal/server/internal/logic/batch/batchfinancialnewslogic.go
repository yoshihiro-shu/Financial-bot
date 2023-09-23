package batch

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchFinancialNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchFinancialNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchFinancialNewsLogic {
	return &BatchFinancialNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchFinancialNewsLogic) BatchFinancialNews(req *types.BatchFinancialNewsReq) (resp *types.BatchFinancialNewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
