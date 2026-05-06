package driverlogic

import (
	"context"

	"substitute-driver/driver/driver"
	"substitute-driver/driver/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNearbyDriversLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNearbyDriversLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNearbyDriversLogic {
	return &GetNearbyDriversLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询附近司机（骨架实现）
func (l *GetNearbyDriversLogic) GetNearbyDrivers(in *driver.GetNearbyDriversReq) (*driver.GetNearbyDriversResp, error) {
	// TODO: 生产环境应使用 Redis GEO 或 MySQL 空间索引优化
	// 当前为骨架实现，返回空列表
	return &driver.GetNearbyDriversResp{
		Drivers: make([]*driver.DriverInfo, 0),
	}, nil
}
