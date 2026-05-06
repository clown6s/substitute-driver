package mapservicelogic

import (
	"context"
	"math"

	"substitute-driver/map/internal/svc"
	"substitute-driver/map/mapservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDistanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDistanceLogic {
	return &GetDistanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDistance 计算两点距离（Haversine 公式）
func (l *GetDistanceLogic) GetDistance(in *mapservice.GetDistanceReq) (*mapservice.GetDistanceResp, error) {
	// TODO: 生产环境接入高德/腾讯地图 API
	// 当前使用 Haversine 公式计算直线距离
	lat1 := in.FromLat * math.Pi / 180
	lat2 := in.ToLat * math.Pi / 180
	dLat := (in.ToLat - in.FromLat) * math.Pi / 180
	dLng := (in.ToLng - in.FromLng) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := 6371 * c                    // 地球半径 6371km
	duration := int64(distance / 40 * 3600) // 假设平均 40km/h

	return &mapservice.GetDistanceResp{
		Distance: distance,
		Duration: duration,
	}, nil
}
