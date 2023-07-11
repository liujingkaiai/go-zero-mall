package svc

import (
	"time"

	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/config"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"golang.org/x/sync/singleflight"
)

const localCahcheExpire = time.Second * 60

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	OperationModel model.ProductOperationModel
	BizRedis       *redis.Redis
	SingleGroup    singleflight.Group
	LocalCache     *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	localCache, err := collection.NewCache(localCahcheExpire)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		ProductModel:   model.NewProductModel(conn, c.CacheRedis),
		CategoryModel:  model.NewCategoryModel(conn, c.CacheRedis),
		OperationModel: model.NewProductOperationModel(conn, c.CacheRedis),
		BizRedis:       redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		LocalCache:     localCache,
	}
}
