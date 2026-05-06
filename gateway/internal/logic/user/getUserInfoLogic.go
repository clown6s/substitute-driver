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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp any, err error) {
	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userclient.GetUserInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return response.Response{
		Code:    0,
		Message: "ok",
		Data: map[string]any{
			"id":        rpcResp.Id,
			"username":  rpcResp.Username,
			"phone":     rpcResp.Phone,
			"avatar":    rpcResp.Avatar,
			"createdAt": rpcResp.CreatedAt,
		},
	}, nil
}
