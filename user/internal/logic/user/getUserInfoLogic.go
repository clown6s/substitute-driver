package userlogic

import (
	"context"
	"fmt"

	"substitute-driver/user/internal/svc"
	"substitute-driver/user/user"

	"substitute-driver/common/errorx"
	"substitute-driver/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserInfo 查询用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	return &user.GetUserInfoResp{
		Id:        userInfo.Id,
		Username:  userInfo.Username,
		Phone:     userInfo.Phone,
		Avatar:    userInfo.Avatar,
		CreatedAt: userInfo.CreatedAt.Unix(),
	}, nil
}
