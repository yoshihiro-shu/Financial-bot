package class

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/logic/class"
	"github.com/yoshihiro-shu/financial-bot/apps/user/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserClassListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := class.NewUserClassListLogic(r.Context(), svcCtx)
		resp, err := l.UserClassList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
