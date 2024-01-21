package product

import (
	"context"
	"grpc-client-product/internal/client"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rizkysr90/my-protobuf/gen/go/personal/productservice/product"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	ProductClient client.ProductServiceClientInterface
}

func NewProductServiceClient(conn *grpc.ClientConn) *ProductServiceClient {
	client := product.NewProductServiceClient(conn)
	return &ProductServiceClient{ProductClient: client}
}

func (p *ProductServiceClient) Create(ctx context.Context,
	product *product.Product) (*product.Response, error) {

	res, err := p.ProductClient.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (p *ProductServiceClient) Update(ctx context.Context,
	product *product.UpdateProduct) (*product.Response, error) {

	res, err := p.ProductClient.Update(ctx, product)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (p *ProductServiceClient) Delete(ctx context.Context,
	product *product.DeleteProduct) (*product.Response, error) {

	res, err := p.ProductClient.Delete(ctx, product)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (p *ProductServiceClient) CreateProducts(ctx context.Context) (
	*product.Response, error) {
	clientStream, err := p.ProductClient.CreateProducts(ctx)
	if err != nil {
		return nil, err
	}
	products := []*product.Product{
		{
			Pid:   1,
			Name:  "Produk 1",
			Stock: 99,
			Price: 990000,
		},
		{
			Pid:   2,
			Name:  "Produk 2",
			Stock: 99,
			Price: 100000,
		},
	}
	for _, product := range products {
		if err := clientStream.Send(product); err != nil {
			return nil, err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return clientStream.CloseAndRecv()
}
func (p *ProductServiceClient) GetListStream(ctx context.Context) ([]*product.Product, error) {
	serverStream, err := p.ProductClient.GetListStream(ctx, new(empty.Empty))
	if err != nil {
		return nil, err
	}
	listProducts := []*product.Product{}
	for {
		product, err := serverStream.Recv()
		if err == io.EOF {
			return listProducts, nil
		}
		if err != nil {
			return nil, err
		}
		listProducts = append(listProducts, product)
	}
}
