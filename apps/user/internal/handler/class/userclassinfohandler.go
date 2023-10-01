package class

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/logic/class"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserClassInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserClassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewUserClassInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserClassInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
