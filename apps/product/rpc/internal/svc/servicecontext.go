package svc

import (
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/config"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ProductModel  model.ProductModel
	CategoryModel model.CategoryModel
	BizRedis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:        c,
		ProductModel:  model.NewProductModel(conn, c.CacheRedis),
		CategoryModel: model.NewCategoryModel(conn, c.CacheRedis),
		BizRedis:      redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
