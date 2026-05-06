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

type ReverseGeocodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewReverseGeocodeLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *ReverseGeocodeLogic {
	return &ReverseGeocodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *ReverseGeocodeLogic) ReverseGeocode(req *types.ReverseGeocodeReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.MapRpc.ReverseGeocode(l.ctx, &mapclient.ReverseGeocodeReq{
		Lng: req.Lng,
		Lat: req.Lat,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"formattedAddress": rpcResp.FormattedAddress,
		},
	}, nil
}
