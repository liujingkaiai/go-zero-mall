package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderCheckLogic {
	return &CreateOrderCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderCheckLogic) CreateOrderCheck(in *rpc.CreateOrderRequest) (*rpc.CreateOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.CreateOrderResponse{}, nil
}
