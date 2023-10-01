package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/logic"
	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/svc"
	"github.com/yoshihiro-shu/financial-bot/apps/notification/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NotificationCryptoCurrencyInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotificationCryptoCurrencyInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNotificationCryptoCurrencyInfoLogic(r.Context(), svcCtx)
		resp, err := l.NotificationCryptoCurrencyInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
