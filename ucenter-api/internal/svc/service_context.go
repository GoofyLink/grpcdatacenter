package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/ucclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc ucclient.Register
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: ucclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
	}
}
