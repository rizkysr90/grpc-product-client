package main

import (
	"context"
	productClient "grpc-client-product/internal/client/product"
	"log"

	"github.com/rizkysr90/my-protobuf/gen/go/personal/productservice/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, err := grpc.Dial("localhost:8888", opts...)
	if err != nil {
		log.Fatalln("failed to dial grpc server")
	}
	productClient := productClient.NewProductServiceClient(grpcConn)
	// res, err := productClient.CreateProducts(context.Background())
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(res)
	res, err := productClient.GetListStream(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	// testCreateNewProduct1(productClient)
	// testCreateNewProduct2(productClient)
	// // testUpdateProductID1(productClient)
	// testDeleteProductID1(productClient)
}

func testCreateNewProduct1(productClient *productClient.ProductServiceClient) {
	data := product.Product{
		Pid:   1,
		Name:  "Tepung Terigu 1",
		Price: 9000,
	}
	res, err := productClient.Create(context.Background(), &data)
	if err != nil {
		log.Fatalln("call_create_product_service : ", err)
	}
	log.Println(res.Status)
}
func testCreateNewProduct2(productClient *productClient.ProductServiceClient) {
	data := product.Product{
		Pid:   2,
		Name:  "Tepung Terigu 2",
		Stock: 10,
		Price: 9000,
	}
	res, err := productClient.Create(context.Background(), &data)
	if err != nil {
		log.Fatalln("call_create_product_service : ", err)
	}
	log.Println(res.Status)
}
func testUpdateProductID1(productClient *productClient.ProductServiceClient) {
	data := product.UpdateProduct{
		Pid: uint32(1),
		UpdatedDataProduct: &product.Product{
			Pid:   uint64(1),
			Name:  "Tepung Terigu 100",
			Stock: 100,
			Price: 1000000,
		},
	}
	res, err := productClient.Update(context.Background(), &data)
	if err != nil {
		log.Fatalln("call_update_product_service : ", err)
	}
	log.Println(res.Status)
}
func testDeleteProductID1(productClient *productClient.ProductServiceClient) {
	data := product.DeleteProduct{
		Pid: uint32(1),
	}
	res, err := productClient.Delete(context.Background(), &data)
	if err != nil {
		log.Fatalln("call_update_product_service : ", err)
	}
	log.Println(res.Status)
}
