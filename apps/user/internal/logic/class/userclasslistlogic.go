package class

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassListLogic {
	return &UserClassListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserClassListLogic) UserClassList() (resp []types.UserClassResp, err error) {
	// todo: add your logic here and delete this line

	return
}
