package batch

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/logic/batch"
	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BatchFinancialNewsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchFinancialNewsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := batch.NewBatchFinancialNewsLogic(r.Context(), svcCtx)
		resp, err := l.BatchFinancialNews(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
