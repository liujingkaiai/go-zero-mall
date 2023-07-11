package logic

import (
	"context"

	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	validStatus          = 1
	operationProductsKey = "operation#products"
)

type OperationProductsLogic struct {
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	productListLogic *ProductListLogic
	logx.Logger
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		productListLogic: NewProductListLogic(ctx, svcCtx),
	}
}

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductRequest) (*product.OperationProductResponse, error) {
	// todo: add your logic here and delete this line
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)
	if ok {
		return &product.OperationProductResponse{
			Products: opProducts.([]*product.ProductItem),
		}, nil
	}
	pos, err := l.svcCtx.OperationModel.OperationProducts(l.ctx, validStatus)
	if err != nil {
		return nil, err
	}

	var pids []int64
	for _, p := range pos {
		pids = append(pids, p.Id)
	}
	products, err := l.productListLogic.productsByIds(l.ctx, pids)
	if err != nil {
		return nil, err
	}

	var pItems []*product.ProductItem

	for _, p := range products {
		pItems = append(pItems, &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		})
	}

	l.svcCtx.LocalCache.Set(operationProductsKey, pItems)
	return &product.OperationProductResponse{
		Products: pItems,
	}, nil
}
