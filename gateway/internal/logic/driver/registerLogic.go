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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *RegisterLogic) Register(req *types.DriverRegisterReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.DriverRpc.Register(l.ctx, &driverclient.RegisterDriverReq{
		Name:      req.Name,
		Phone:     req.Phone,
		LicenseNo: req.LicenseNo,
		CarNo:     req.CarNo,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"id": rpcResp.Id,
		},
	}, nil
}
