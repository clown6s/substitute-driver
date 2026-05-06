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

type UpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UpdateStatusLogic) UpdateStatus(req *types.UpdateDriverStatusReq) (resp any, err error) {
	_, err = l.svcCtx.DriverRpc.UpdateStatus(l.ctx, &driverclient.UpdateDriverStatusReq{
		Id:     int64(req.Id),
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
	}, nil
}
