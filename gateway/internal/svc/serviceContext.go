// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"substitute-driver/gateway/internal/config"

	driverclient "substitute-driver/driver/client/driver"
	mapclient "substitute-driver/map/client/mapservice"
	userclient "substitute-driver/user/client/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc   userclient.User
	DriverRpc driverclient.Driver
	MapRpc    mapclient.Mapservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		DriverRpc: driverclient.NewDriver(zrpc.MustNewClient(c.DriverRpc)),
		MapRpc:    mapclient.NewMapservice(zrpc.MustNewClient(c.MapRpc)),
	}
}
