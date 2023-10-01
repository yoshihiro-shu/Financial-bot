package role

import (
	"context"

	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleAddLogic {
	return &UserRoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleAddLogic) UserRoleAdd(req *types.UserRoleAddReq) (resp *types.UserRoleAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
