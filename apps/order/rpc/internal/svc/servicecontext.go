package svc

import (
	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/internal/config"
	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrdersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrdersModel(conn, c.CacheRedis),
	}
}
