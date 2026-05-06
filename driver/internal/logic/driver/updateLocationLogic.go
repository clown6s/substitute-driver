package driverlogic

import (
	"context"
	"fmt"

	"substitute-driver/driver/driver"
	"substitute-driver/driver/internal/svc"

	"substitute-driver/common/errorx"
	"substitute-driver/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLocationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLocationLogic {
	return &UpdateLocationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新司机位置
func (l *UpdateLocationLogic) UpdateLocation(in *driver.UpdateDriverLocationReq) (*driver.UpdateDriverLocationResp, error) {
	driverInfo, err := l.svcCtx.DriverModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.ErrDriverNotFound
		}
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	driverInfo.Longitude = in.Longitude
	driverInfo.Latitude = in.Latitude

	err = l.svcCtx.DriverModel.Update(l.ctx, driverInfo)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	return &driver.UpdateDriverLocationResp{}, nil
}
