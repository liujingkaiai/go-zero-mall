package main

import (
	"flag"
	"fmt"

	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/internal/config"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/internal/server"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rpc/seckill"
	"github.com/liujingkaiai/go-zero-mall/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		seckill.RegisterSeckillServer(grpcServer, server.NewSeckillServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
