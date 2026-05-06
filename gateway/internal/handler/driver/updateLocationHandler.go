package driver

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"substitute-driver/gateway/internal/logic/driver"
	"substitute-driver/gateway/internal/svc"
	"substitute-driver/gateway/internal/types"
)

func UpdateLocationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDriverLocationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := driver.NewUpdateLocationLogic(r.Context(), svcCtx, r)
		resp, err := l.UpdateLocation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
