package batch

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCryptoCurrencyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchCryptoCurrencyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCryptoCurrencyInfoLogic {
	return &BatchCryptoCurrencyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchCryptoCurrencyInfoLogic) BatchCryptoCurrencyInfo(req *types.BatchCryptoCurrencyReq) (resp *types.BatchCryptoCurrencyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
