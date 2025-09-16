package service

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	catalogv1 "github.com/materials-resources/microservices-proto/pkg/catalog"
	"github.com/materials-resources/microservices-proto/pkg/catalog/catalogconnect"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/internal/domain"
	"github.com/materials-resources/store-api/internal/oas"
	"google.golang.org/protobuf/proto"
)

type CatalogService struct {
	Client catalogconnect.CatalogServiceClient
}

func NewCatalogService(a *app.App) *CatalogService {
	otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
	if err != nil {
		a.Logger.Fatal().Str("service", "catalog").Err(err).Msg("could not create otel interceptor")
	}
	return &CatalogService{
		Client: catalogconnect.NewCatalogServiceClient(http.DefaultClient,
			a.Config.Services.CatalogUrl,
			connect.WithInterceptors(otelInterceptor),
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
		IsActive: pbRes.Msg.GetProduct().GetIsActive(),
		HasStock: pbRes.Msg.GetProduct().GetHasStock(),
	}

	return &product, nil

}
