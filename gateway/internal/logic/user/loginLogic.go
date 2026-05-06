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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"id":    rpcResp.Id,
			"token": rpcResp.Token,
		},
	}, nil
}
