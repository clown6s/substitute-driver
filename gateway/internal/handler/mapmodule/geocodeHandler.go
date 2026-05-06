package mapmodule

import (
	"net/http"

	"substitute-driver/gateway/internal/logic/mapmodule"
	"substitute-driver/gateway/internal/svc"
	"substitute-driver/gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GeocodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeocodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mapmodule.NewGeocodeLogic(r.Context(), svcCtx, r)
		resp, err := l.Geocode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
