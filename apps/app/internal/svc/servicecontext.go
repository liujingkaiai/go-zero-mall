package svc

import (
	"app/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
