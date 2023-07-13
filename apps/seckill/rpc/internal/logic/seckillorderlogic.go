package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/product"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/seckill"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	limitPeriod       = 10
	limitQuota        = 1
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60
)

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

type SeckillOrderLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	logx.Logger
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	return &SeckillOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
		limiter:    limit.NewPeriodLimit(limitPeriod, limitQuota, svcCtx.BizRedis, seckillUserPrefix),
	}
}

func (l *SeckillOrderLogic) SeckillOrder(in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {
	// todo: add your logic here and delete this line
	code, _ := l.limiter.Take(strconv.FormatInt(in.UserId, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}
	//调用product.rpc获取商品实例
	p, err := l.svcCtx.ProductRPC.Product(l.ctx, &product.ProductItemRequest{
		ProductId: in.ProductId,
	})
	if err != nil {
		return nil, err
	}
	//判断商品库存
	if p.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}

	kd, err := json.Marshal(&KafkaData{Uid: in.UserId, Pid: p.ProductId})
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
		return nil, err
	}
	return &seckill.SeckillOrderResponse{}, nil
}
