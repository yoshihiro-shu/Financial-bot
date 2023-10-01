package batch

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/logic/batch"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchFinancialNewsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := batch.NewHealthCheckLogic(r.Context(), svcCtx)
		resp, err := l.HealthCheck(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
