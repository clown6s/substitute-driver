package mapservicelogic

import (
	"context"

	"substitute-driver/map/internal/svc"
	"substitute-driver/map/mapservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRouteLogic {
	return &GetRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRoute 路线规划
func (l *GetRouteLogic) GetRoute(in *mapservice.GetRouteReq) (*mapservice.GetRouteResp, error) {
	// TODO: 生产环境接入高德/腾讯地图路线规划 API
	// 当前返回起终点直线路径
	return &mapservice.GetRouteResp{
		Distance: 0,
		Duration: 0,
		Polyline: []*mapservice.Point{
			{Lng: in.FromLng, Lat: in.FromLat},
			{Lng: in.ToLng, Lat: in.ToLat},
		},
	}, nil
}
