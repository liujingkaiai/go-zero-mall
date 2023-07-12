package rmq

import (
	"flag"

	"github.com/liujingkaiai/go-zero-mall/apps/seckill/rmq/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

}
