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

type GetDriverInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDriverInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDriverInfoLogic {
	return &GetDriverInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询司机信息
func (l *GetDriverInfoLogic) GetDriverInfo(in *driver.GetDriverInfoReq) (*driver.GetDriverInfoResp, error) {
	driverInfo, err := l.svcCtx.DriverModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.ErrDriverNotFound
		}
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	return &driver.GetDriverInfoResp{
		Driver: &driver.DriverInfo{
			Id:        driverInfo.Id,
			Name:      driverInfo.Name,
			Phone:     driverInfo.Phone,
			LicenseNo: driverInfo.LicenseNo,
			CarNo:     driverInfo.CarNo,
			Status:    int32(driverInfo.Status),
			Longitude: driverInfo.Longitude,
			Latitude:  driverInfo.Latitude,
		},
	}, nil
}
