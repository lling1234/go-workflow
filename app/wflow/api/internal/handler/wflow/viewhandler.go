package wflow

import (
	"net/http"

	"go-wflow/app/wflow/api/internal/logic/wflow"
	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/rest/httpx"
)

func ViewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteProcInstReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := wflow.NewViewLogic(r.Context(), svcCtx)
		resp, err := l.View(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
