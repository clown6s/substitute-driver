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

type GetDistanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetDistanceLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetDistanceLogic {
	return &GetDistanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetDistanceLogic) GetDistance(req *types.GetDistanceReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.MapRpc.GetDistance(l.ctx, &mapclient.GetDistanceReq{
		FromLng: req.FromLng,
		FromLat: req.FromLat,
		ToLng:   req.ToLng,
		ToLat:   req.ToLat,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"distance": rpcResp.Distance,
			"duration": rpcResp.Duration,
		},
	}, nil
}
