package act

import (
	"net/http"

	"act/api/internal/logic/act"
	"act/api/internal/svc"
	"act/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindByTargetIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindByTargetId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := act.NewFindByTargetIdLogic(r.Context(), svcCtx)
		resp, err := l.FindByTargetId(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
