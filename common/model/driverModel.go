package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DriverModel = (*customDriverModel)(nil)

type (
	// DriverModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDriverModel.
	DriverModel interface {
		driverModel
	}

	customDriverModel struct {
		*defaultDriverModel
	}
)

// NewDriverModel returns a model for the database table.
func NewDriverModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DriverModel {
	return &customDriverModel{
		defaultDriverModel: newDriverModel(conn, c, opts...),
	}
}
