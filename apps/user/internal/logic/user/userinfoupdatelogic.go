package user

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateReq) (resp *types.UserInfoUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
