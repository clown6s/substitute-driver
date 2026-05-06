// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package mapmodule

import (
	"context"
	"net/http"

	"substitute-driver/common/response"
	"substitute-driver/gateway/internal/svc"
	"substitute-driver/gateway/internal/types"

	mapclient "substitute-driver/map/client/mapservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeocodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGeocodeLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GeocodeLogic {
	return &GeocodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GeocodeLogic) Geocode(req *types.GeocodeReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.MapRpc.Geocode(l.ctx, &mapclient.GeocodeReq{
		Address: req.Address,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"lng":              rpcResp.Lng,
			"lat":              rpcResp.Lat,
			"formattedAddress": rpcResp.FormattedAddress,
		},
	}, nil
}
