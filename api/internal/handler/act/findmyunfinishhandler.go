package act

import (
	"net/http"

	"act/api/internal/logic/act"
	"act/api/internal/svc"
	"act/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindMyUnFinishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := act.NewFindMyUnFinishLogic(r.Context(), svcCtx)
		resp, err := l.FindMyUnFinish(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
