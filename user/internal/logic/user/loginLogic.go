package userlogic

import (
	"context"
	"fmt"
	"time"

	"substitute-driver/user/internal/svc"
	"substitute-driver/user/user"

	"substitute-driver/common/errorx"
	"substitute-driver/common/model"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", errorx.ErrDBError, err)
	}

	// 用 bcrypt 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.Password)); err != nil {
		return nil, errorx.ErrPasswordWrong
	}

	// 签发 JWT
	token, err := generateToken(l.svcCtx.Config.JWTConfig.AccessSecret, userInfo.Id, l.svcCtx.Config.JWTConfig.AccessExpire)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorx.ErrServerInternal, err)
	}

	return &user.LoginResp{
		Id:    userInfo.Id,
		Token: token,
	}, nil
}

// generateToken 签发 JWT token
func generateToken(secret string, userId int64, expireSeconds int64) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    now.Add(time.Duration(expireSeconds) * time.Second).Unix(),
		"iat":    now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
