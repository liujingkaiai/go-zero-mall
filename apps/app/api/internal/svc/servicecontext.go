package svc

import (
	"github.com/liujingkaiai/go-zero-mall/apps/app/api/internal/config"
	order "github.com/liujingkaiai/go-zero-mall/apps/order/rpc/orderclient"
	product "github.com/liujingkaiai/go-zero-mall/apps/product/rpc/productclient"

	reply "github.com/liujingkaiai/go-zero-mall/apps/reply/rpc/replyclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
	ReplyRPC   reply.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		ReplyRPC:   reply.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
	}
}
