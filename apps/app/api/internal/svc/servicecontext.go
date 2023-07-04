package svc

import (
	"github.com/liujingkaiai/go-zero-mall/apps/app/api/internal/config"
	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/order"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/product"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRPC order.Order
	Product  product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
