package errorx

import "google.golang.org/grpc/status"

// 统一错误码
var (
	ErrUserAlreadyExists = status.Error(10001, "用户已存在")
	ErrUserNotFound      = status.Error(10002, "用户不存在")
	ErrPasswordWrong     = status.Error(10003, "密码错误")
	ErrInvalidParams     = status.Error(10004, "参数错误")
	ErrDBError           = status.Error(10005, "数据库错误")
	ErrServerInternal    = status.Error(10006, "服务内部错误")

	ErrDriverAlreadyExists = status.Error(20001, "司机已存在")
	ErrDriverNotFound      = status.Error(20002, "司机不存在")
	ErrDriverOffline       = status.Error(20003, "司机不在线")

	ErrMapServiceError = status.Error(30001, "地图服务异常")
)
