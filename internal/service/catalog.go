package service

import (
	"connectrpc.com/connect"
	"github.com/materials-resources/customer-api/internal/grpc-client/catalog/catalogconnect"
	"net/http"
)

type Catalog struct {
	Client catalogconnect.CatalogServiceClient
}

func NewCatalogService() *Catalog {
	return &Catalog{
		Client: catalogconnect.NewCatalogServiceClient(http.DefaultClient,
			"http://localhost:50058",
			connect.WithGRPC()),
	}
}
