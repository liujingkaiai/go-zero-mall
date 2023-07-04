package logic

import (
	"context"
	"strconv"
	"strings"

	"github.com/liujingkaiai/go-zero-mall/apps/app/api/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/app/api/internal/types"
	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/order"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderListResponse, err error) {
	// todo: add your logic here and delete this line
	orderRet, err := l.svcCtx.OrderRPC.Orders(l.ctx, &order.OrdersRequest{UserId: req.UID})
	if err != nil {
		return nil, err
	}
	var pids []string
	for _, o := range orderRet.Orders {
		pids = append(pids, strconv.Itoa(int(o.Proid)))
	}
	productRet, err := l.svcCtx.ProductRPC.Products(l.ctx, &product.ProductRequest{ProductIds: strings.Join(pids, ",")})
	if err != nil {
		return nil, err
	}
	var orders []*types.Order
	for _, o := range orderRet.Orders {
		if p := productRet.Products[o.Proid]; p != nil {
			orders = append(orders, &types.Order{
				OrderID:     o.Orderid,
				ProductName: p.Name,
			})
		}
	}
	return &types.OrderListResponse{Orders: orders}, nil
}
