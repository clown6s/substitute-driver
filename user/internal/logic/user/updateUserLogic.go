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

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	// 先查询用户是否存在
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	// 更新字段（保留原密码和不可修改字段）
	updatedUser := *userInfo
	if in.Username != "" {
		updatedUser.Username = in.Username
	}
	if in.Phone != "" {
		updatedUser.Phone = in.Phone
	}
	if in.Avatar != "" {
		updatedUser.Avatar = in.Avatar
	}

	err = l.svcCtx.UserModel.Update(l.ctx, &updatedUser)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	return &user.UpdateUserResp{}, nil
}
