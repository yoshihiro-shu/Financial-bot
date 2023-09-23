package batch

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/logic/batch"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/batch/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BatchStockInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchStockInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := batch.NewBatchStockInfoLogic(r.Context(), svcCtx)
		resp, err := l.BatchStockInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
