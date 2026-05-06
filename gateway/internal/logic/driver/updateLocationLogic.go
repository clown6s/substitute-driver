// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package driver

import (
	"context"
	"net/http"

	"substitute-driver/common/response"
	"substitute-driver/gateway/internal/svc"
	"substitute-driver/gateway/internal/types"

	driverclient "substitute-driver/driver/client/driver"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateLocationLogic {
	return &UpdateLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UpdateLocationLogic) UpdateLocation(req *types.UpdateDriverLocationReq) (resp any, err error) {
	_, err = l.svcCtx.DriverRpc.UpdateLocation(l.ctx, &driverclient.UpdateDriverLocationReq{
		Id:        req.Id,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
	}, nil
}
