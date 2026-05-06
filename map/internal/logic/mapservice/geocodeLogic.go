package mapservicelogic

import (
	"context"

	"substitute-driver/map/internal/svc"
	"substitute-driver/map/mapservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeocodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeocodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeocodeLogic {
	return &GeocodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Geocode 地理编码（地址转坐标）
func (l *GeocodeLogic) Geocode(in *mapservice.GeocodeReq) (*mapservice.GeocodeResp, error) {
	// TODO: 接入高德/腾讯地图地理编码 API
	return &mapservice.GeocodeResp{
		Lng:              0,
		Lat:              0,
		FormattedAddress: in.Address,
	}, nil
}
