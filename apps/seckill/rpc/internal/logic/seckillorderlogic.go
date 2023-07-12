package logic

import (
	"context"

	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	return &SeckillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SeckillOrderLogic) SeckillOrder(in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &seckill.SeckillOrderResponse{}, nil
}
