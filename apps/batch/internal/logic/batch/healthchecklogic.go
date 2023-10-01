package batch

import (
	"context"
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck(req *types.BatchFinancialNewsReq) (resp *types.BatchFinancialNewsResp, err error) {
	// todo: add your logic here and delete this line

	return &types.BatchFinancialNewsResp{
		BaseResp: &types.BaseResp{
			Code: http.StatusOK,
			Msg:  "OK",
		},
	}, nil
}
