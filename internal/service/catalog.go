package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/internal/domain"
	"github.com/materials-resources/store-api/internal/oas"
	catalogv1 "github.com/materials-resources/store-api/internal/proto/catalog"
	"github.com/materials-resources/store-api/internal/proto/catalog/catalogconnect"
	"google.golang.org/protobuf/proto"
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

func (s *CatalogService) GetProduct(ctx context.Context, params oas.GetProductParams) (*domain.Product, error) {
	pbReq := catalogv1.GetProductRequest_builder{Id: proto.String(params.ID)}.Build()

	pbRes, err := s.Client.GetProduct(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	product := domain.Product{
		Id:               pbRes.Msg.GetProduct().GetId(),
		Sn:               pbRes.Msg.GetProduct().GetSn(),
		Name:             pbRes.Msg.GetProduct().GetName(),
		Description:      pbRes.Msg.GetProduct().GetDescription(),
		ProductGroupId:   pbRes.Msg.GetProduct().GetProductGroupSn(),
		ProductGroupName: pbRes.Msg.GetProduct().GetProductGroupName(),
		SalesUnitOfMeasurement: domain.UnitOfMeasurement{
			Id:               pbRes.Msg.GetProduct().GetSalesUnitOfMeasurement().GetId(),
			ConversionFactor: pbRes.Msg.GetProduct().GetSalesUnitOfMeasurement().GetConversionFactor(),
		},
	}

	return &product, nil

}
