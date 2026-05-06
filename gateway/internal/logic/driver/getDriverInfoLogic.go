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

type GetDriverInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetDriverInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetDriverInfoLogic {
	return &GetDriverInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetDriverInfoLogic) GetDriverInfo(req *types.GetDriverInfoReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.DriverRpc.GetDriverInfo(l.ctx, &driverclient.GetDriverInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"id":        rpcResp.Driver.Id,
			"name":      rpcResp.Driver.Name,
			"phone":     rpcResp.Driver.Phone,
			"licenseNo": rpcResp.Driver.LicenseNo,
			"carNo":     rpcResp.Driver.CarNo,
			"status":    rpcResp.Driver.Status,
			"longitude": rpcResp.Driver.Longitude,
			"latitude":  rpcResp.Driver.Latitude,
		},
	}, nil
}
