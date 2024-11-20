package service

import (
	"customer-api/client/microservices/proto/catalog/v1"
	"google.golang.org/grpc"
)

type Catalog struct {
	Client catalog.CatalogServiceClient
}

func NewCatalogService() *Catalog {
	conn, err := grpc.Dial("localhost:50058", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &Catalog{
		Client: catalog.NewCatalogServiceClient(conn),
	}
}
