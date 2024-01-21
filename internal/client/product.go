package client

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rizkysr90/my-protobuf/gen/go/personal/productservice/product"
	"google.golang.org/grpc"
)

type ProductServiceClientInterface interface {
	// create new product
	Create(ctx context.Context, in *product.Product, opts ...grpc.CallOption) (*product.Response, error)
	// get list of product data
	GetList(ctx context.Context, in *product.ListProduct, opts ...grpc.CallOption) (*product.Response, error)
	// update product data
	Update(ctx context.Context, in *product.UpdateProduct, opts ...grpc.CallOption) (*product.Response, error)
	// delete product data
	Delete(ctx context.Context, in *product.DeleteProduct, opts ...grpc.CallOption) (*product.Response, error)
	// create new product streams
	CreateProducts(ctx context.Context, opts ...grpc.CallOption) (product.ProductService_CreateProductsClient, error)
	// get all products streams
	GetListStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (product.ProductService_GetListStreamClient, error)
}
