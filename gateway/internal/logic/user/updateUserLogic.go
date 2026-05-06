// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"
	"net/http"

	"substitute-driver/common/response"
	"substitute-driver/gateway/internal/svc"
	"substitute-driver/gateway/internal/types"

	userclient "substitute-driver/user/client/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp any, err error) {
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &userclient.UpdateUserReq{
		Id:       req.Id,
		Username: req.Username,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
	}, nil
}
