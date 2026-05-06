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

type GetRouteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetRouteLogic {
	return &GetRouteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetRouteLogic) GetRoute(req *types.GetRouteReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.MapRpc.GetRoute(l.ctx, &mapclient.GetRouteReq{
		FromLng: req.FromLng,
		FromLat: req.FromLat,
		ToLng:   req.ToLng,
		ToLat:   req.ToLat,
	})
	if err != nil {
		return nil, err
	}

	polyline := make([]map[string]any, 0, len(rpcResp.Polyline))
	for _, p := range rpcResp.Polyline {
		polyline = append(polyline, map[string]any{
			"lng": p.Lng,
			"lat": p.Lat,
		})
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"distance": rpcResp.Distance,
			"duration": rpcResp.Duration,
			"polyline": polyline,
		},
	}, nil
}
