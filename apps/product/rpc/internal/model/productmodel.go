package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		CategoryProducts(ctx context.Context, ctime string, cateid, limit int64) ([]*Product, error)
		UpdateProductStock(ctx context.Context, pid, num int64) error
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

func (c customProductModel) CategoryProducts(ctx context.Context, ctime string, cateid, limit int64) ([]*Product, error) {
	var products []*Product
	err := c.QueryRowsNoCacheCtx(ctx, &products, fmt.Sprintf("select %s from %s where cateid=? and status=1 and create_time<? order by create_time desc limit ?", productRows, c.table), cateid, ctime, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// 修改商品库存
func (m *customProductModel) UpdateProductStock(ctx context.Context, pid, num int64) error {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, pid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET stock = stock - ? WHERE id = ?", m.table), num, pid)
	}, productProductIdKey)
	return err
}
