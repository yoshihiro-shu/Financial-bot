package user

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/logic/user"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserListLogic(r.Context(), svcCtx)
		resp, err := l.UserList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
