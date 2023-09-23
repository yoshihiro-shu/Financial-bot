// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	batch "github.com/yoshihiro-shu/financial-bot/internal/server/internal/handler/batch"
	"github.com/yoshihiro-shu/financial-bot/internal/server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/batch/financial-news",
				Handler: batch.BatchFinancialNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/batch/stock-info",
				Handler: batch.BatchStockInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/batch/cypto-currency-info",
				Handler: batch.BatchCryptoCurrencyInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
