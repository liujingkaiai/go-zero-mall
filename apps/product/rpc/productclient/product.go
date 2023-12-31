// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productclient

import (
	"context"

	"github.com/liujingkaiai/go-zero-mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	OperationProductsRequest   = product.OperationProductsRequest
	OperationProductsResponse  = product.OperationProductsResponse
	ProductItem                = product.ProductItem
	ProductItemRequest         = product.ProductItemRequest
	ProductListRequest         = product.ProductListRequest
	ProductListResponse        = product.ProductListResponse
	ProductRequest             = product.ProductRequest
	ProductResponse            = product.ProductResponse
	UpdateProductStockRequest  = product.UpdateProductStockRequest
	UpdateProductStockResponse = product.UpdateProductStockResponse

	Product interface {
		Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error)
		Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
		ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error)
		OperationProducts(ctx context.Context, in *OperationProductsRequest, opts ...grpc.CallOption) (*OperationProductsResponse, error)
		UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Product(ctx, in, opts...)
}

func (m *defaultProduct) Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Products(ctx, in, opts...)
}

func (m *defaultProduct) ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}

func (m *defaultProduct) OperationProducts(ctx context.Context, in *OperationProductsRequest, opts ...grpc.CallOption) (*OperationProductsResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.OperationProducts(ctx, in, opts...)
}

func (m *defaultProduct) UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateProductStock(ctx, in, opts...)
}
