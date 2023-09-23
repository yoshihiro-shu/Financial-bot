package batch

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchStockInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchStockInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchStockInfoLogic {
	return &BatchStockInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchStockInfoLogic) BatchStockInfo(req *types.BatchStockInfoReq) (resp *types.BatchStockInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
