### 目录结构 

- app - BFF服务
- cart - 购物车服务
- order - 订单服务
- pay - 支付服务
- product - 商品服务
- recommend - 推荐服务
- reply - 评论服务
- user - 账号服务



### 调试 

#### 安装grpurl 

-go install github.com/fullstorydev/grpcurl/cmd/grpcurl 

#### 开设反射服务 

```
在/etc/xxx.yaml  

设置 Mode: dev
```
启动服务，通过如下命令查询服务，服务提供的方法，可以看到当前提供了Product获取商品详情接口和Products批量获取商品详情接口

```shell

~ grpcurl -plaintext 127.0.0.1:8081 list

grpc.health.v1.Health
grpc.reflection.v1alpha.ServerReflection
product.Product

~ grpcurl -plaintext 127.0.0.1:8081 list product.Product
product.Product.Product
product.Product.Products
```

通过grpcurl工具来调用Product接口查询id为1的商品数据

```shell 
~ grpcurl -plaintext -d '{"product_id": 1}' 127.0.0.1:8081 product.Product.Product
{
  "productId": "1",
  "name": "夹克1"
}
```