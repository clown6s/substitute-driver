package svc

import (
	"substitute-driver/driver/internal/config"

	"substitute-driver/common/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	DriverModel model.DriverModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		DriverModel: model.NewDriverModel(conn, c.CacheRedis),
	}
}
