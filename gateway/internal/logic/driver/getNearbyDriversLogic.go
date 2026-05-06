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

type GetNearbyDriversLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetNearbyDriversLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetNearbyDriversLogic {
	return &GetNearbyDriversLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetNearbyDriversLogic) GetNearbyDrivers(req *types.GetNearbyDriversReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.DriverRpc.GetNearbyDrivers(l.ctx, &driverclient.GetNearbyDriversReq{
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		Radius:    req.Radius,
	})
	if err != nil {
		return nil, err
	}

	drivers := make([]map[string]any, 0, len(rpcResp.Drivers))
	for _, d := range rpcResp.Drivers {
		drivers = append(drivers, map[string]any{
			"id":        d.Id,
			"name":      d.Name,
			"phone":     d.Phone,
			"licenseNo": d.LicenseNo,
			"carNo":     d.CarNo,
			"status":    d.Status,
			"longitude": d.Longitude,
			"latitude":  d.Latitude,
		})
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"drivers": drivers,
		},
	}, nil
}
