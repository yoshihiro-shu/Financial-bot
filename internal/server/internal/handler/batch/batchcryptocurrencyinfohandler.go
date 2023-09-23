package batch

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/logic/batch"
	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BatchCryptoCurrencyInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchCryptoCurrencyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := batch.NewBatchCryptoCurrencyInfoLogic(r.Context(), svcCtx)
		resp, err := l.BatchCryptoCurrencyInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
