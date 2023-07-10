package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/model"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	readLimit = 10 * time.Millisecond
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	singleKey := fmt.Sprintf("product:%d", in.ProductId)
	v, err, _ := l.svcCtx.SingleGroup.Do(singleKey, func() (interface{}, error) {
		go func() {
			time.Sleep(readLimit)
			l.svcCtx.SingleGroup.Forget(singleKey)
		}()
		return l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	})

	if err != nil {
		return nil, err
	}

	p, ok := v.(*model.Product)
	if !ok {
		return nil, errors.New("product model type err ")
	}

	return &product.ProductItem{
		ProductId: p.Id,
		Name:      p.Name,
	}, nil
}
