package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrdersModel = (*customOrdersModel)(nil)

type (
	// OrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrdersModel.
	OrdersModel interface {
		ordersModel
		CreateOrder(ctx context.Context, oid string, uid, pid int64) error
	}

	customOrdersModel struct {
		*defaultOrdersModel
	}
)

// NewOrdersModel returns a model for the database table.
func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrdersModel {
	return &customOrdersModel{
		defaultOrdersModel: newOrdersModel(conn, c, opts...),
	}
}

func (m *customOrdersModel) CreateOrder(ctx context.Context, oid string, uid, pid int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		err := conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
			_, err := session.ExecCtx(ctx, "INSERT INTO orders(id, userid) VALUES(?,?)", oid, uid)
			if err != nil {
				return err
			}
			_, err = session.ExecCtx(ctx, "INSERT INTO orderitem(orderid, userid, proid) VALUES(?,?,?)", "", uid, pid)
			return err
		})
		return nil, err
	})

	return err
}
