package driverlogic

import (
	"context"
	"fmt"

	"substitute-driver/driver/driver"
	"substitute-driver/driver/internal/svc"

	"substitute-driver/common/errorx"
	"substitute-driver/common/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 司机注册
func (l *RegisterLogic) Register(in *driver.RegisterDriverReq) (*driver.RegisterDriverResp, error) {
	// 检查手机号是否已存在
	_, err := l.svcCtx.DriverModel.FindOneByPhone(l.ctx, in.Phone)
	if err == nil {
		return nil, errorx.ErrDriverAlreadyExists
	}
	if err != sqlx.ErrNotFound {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	result, err := l.svcCtx.DriverModel.Insert(l.ctx, &model.Driver{
		Name:      in.Name,
		Phone:     in.Phone,
		LicenseNo: in.LicenseNo,
		CarNo:     in.CarNo,
		Status:    0, // 默认休息状态
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	id, _ := result.LastInsertId()

	return &driver.RegisterDriverResp{Id: id}, nil
}
