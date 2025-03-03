package service

import (
	"connectrpc.com/connect"
	"context"
	catalogv1 "github.com/materials-resources/store-api/internal/grpc-client/catalog"
	"github.com/materials-resources/store-api/internal/grpc-client/catalog/catalogconnect"
	"github.com/materials-resources/store-api/internal/oas"
	"net/http"
)

type CatalogService struct {
	Client catalogconnect.CatalogServiceClient
}

func NewCatalogService() *CatalogService {
	return &CatalogService{
		Client: catalogconnect.NewCatalogServiceClient(http.DefaultClient,
			"http://localhost:8082",
			connect.WithGRPC()),
	}
}

func (s *CatalogService) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	pbReq := &catalogv1.GetProductRequest{Id: params.ID}

	pbRes, err := s.Client.GetProduct(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	response := oas.GetProductOK{
		Product: oas.Product{
			ID:               pbRes.Msg.GetProduct().GetId(),
			Sn:               pbRes.Msg.GetProduct().GetSn(),
			Name:             pbRes.Msg.GetProduct().GetName(),
			ProductGroupSn:   pbRes.Msg.GetProduct().GetProductGroupSn(),
			ProductGroupName: pbRes.Msg.GetProduct().GetProductGroupName(),
			Description:      oas.OptString{Value: pbRes.Msg.GetProduct().GetDescription(), Set: true},
			ImageURL:         oas.OptString{},
		},
	}

	return &response, nil

}
