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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
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
