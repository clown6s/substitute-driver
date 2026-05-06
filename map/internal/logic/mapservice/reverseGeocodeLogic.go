package mapservicelogic

import (
	"context"

	"substitute-driver/map/internal/svc"
	"substitute-driver/map/mapservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReverseGeocodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReverseGeocodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReverseGeocodeLogic {
	return &ReverseGeocodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 逆地理编码（坐标转地址）
func (l *ReverseGeocodeLogic) ReverseGeocode(in *mapservice.ReverseGeocodeReq) (*mapservice.ReverseGeocodeResp, error) {
	// TODO: 接入高德/腾讯地图逆地理编码 API
	return &mapservice.ReverseGeocodeResp{
		FormattedAddress: "mock address",
	}, nil
}
