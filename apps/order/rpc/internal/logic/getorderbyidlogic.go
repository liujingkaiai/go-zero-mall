package logic

import (
	"context"

	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *order.GetOrderByIdReq) (*order.GetOrderByIdResp, error) {
	// todo: add your logic here and delete this line

	return &order.GetOrderByIdResp{}, nil
}
