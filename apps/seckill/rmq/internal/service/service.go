package service

import (
	"github.com/liujingkaiai/go-zero-mall/apps/order/rpc/orderclient"
	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/productclient"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rmq/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Service struct {
	c          config.Config
	ProductRPC productclient.Product
	OrderPRC   orderclient.Order
}

type KafkaData struct {
	Uid int64 `json:"uid"` //user_id
	Pid int64 `json:"pid"` //product_id
}

func NewService(c config.Config) *Service {
	return &Service{
		c:          c,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderPRC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderPRC)),
	}
}

func (s *Service) Consume(_ string, value string) (err error) {
	logx.Info("Consume value:%s \n", value)
	// var data KafkaData
	// if err = json.Unmarshal([]byte(value), &data); err != nil {
	// 	return
	// }

	// p, err := s.ProductRPC.Product(context.Background(), &product.ProductItemRequest{ProductId: data.Pid})
	// if err != nil {
	// 	return err
	// }

	return
}
