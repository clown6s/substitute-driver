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

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新司机状态
func (l *UpdateStatusLogic) UpdateStatus(in *driver.UpdateDriverStatusReq) (*driver.UpdateDriverStatusResp, error) {
	driverInfo, err := l.svcCtx.DriverModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.ErrDriverNotFound
		}
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	driverInfo.Status = int64(in.Status)

	err = l.svcCtx.DriverModel.Update(l.ctx, driverInfo)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	return &driver.UpdateDriverStatusResp{}, nil
}
