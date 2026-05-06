package userlogic

import (
	"context"
	"fmt"

	"substitute-driver/user/internal/svc"
	"substitute-driver/user/user"

	"substitute-driver/common/errorx"
	"substitute-driver/common/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 检查用户名是否已存在
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, errorx.ErrUserAlreadyExists
	}
	if err != sqlx.ErrNotFound {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	// 检查手机号是否已存在
	_, err = l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err == nil {
		return nil, errorx.ErrUserAlreadyExists
	}
	if err != sqlx.ErrNotFound {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	// 密码 bcrypt 哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: string(hash),
		Phone:    in.Phone,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	id, _ := result.LastInsertId()

	return &user.RegisterResp{Id: id}, nil
}
