package mapmodule

import (
	"net/http"

	"substitute-driver/gateway/internal/logic/mapmodule"
	"substitute-driver/gateway/internal/svc"
	"substitute-driver/gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRouteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRouteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mapmodule.NewGetRouteLogic(r.Context(), svcCtx, r)
		resp, err := l.GetRoute(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
