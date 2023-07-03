package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMLogic {
	return &CreateOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderDTMLogic) CreateOrderDTM(in *rpc.AddOrderReq) (*rpc.AddOrderResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.AddOrderResp{}, nil
}
