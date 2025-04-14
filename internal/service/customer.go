package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/internal/domain"
	customerv1 "github.com/materials-resources/store-api/internal/proto/customer"
	"github.com/materials-resources/store-api/internal/proto/customer/customerconnect"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type CustomerService struct {
	Client customerconnect.CustomerServiceClient
}

func NewCustomerService(a *app.App) *CustomerService {
	otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
	if err != nil {
		a.Logger.Fatal().Str("service", "customer").Err(err).Msg("could not create otel interceptor")
	}
	return &CustomerService{
		Client: customerconnect.NewCustomerServiceClient(http.DefaultClient, a.Config.Services.CustomerUrl, connect.WithInterceptors(otelInterceptor),
			connect.WithGRPC()),
	}
}

func (s *CustomerService) GetRecentPurchases(ctx context.Context, page int, pageSize int,
	branchId string) ([]*domain.PurchaseSummary, int, error) {
	pbReq := customerv1.GetRecentPurchasesByBranchRequest_builder{
		Id:       proto.String(branchId),
		PageSize: proto.Int32(int32(pageSize)),
		Page:     proto.Int32(int32(page)),
	}.Build()

	pbRes, err := s.Client.GetRecentPurchasesByBranch(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, 0, err
	}

	var purchases []*domain.PurchaseSummary
	for _, item := range pbRes.Msg.GetItems() {
		purchases = append(purchases, &domain.PurchaseSummary{
			ProductId:          item.GetProductId(),
			ProductSn:          item.GetProductSn(),
			ProductName:        item.GetProductName(),
			ProductDescription: item.GetProductDescription(),
			OrderedQuantity:    item.GetOrderedQuantity(),
			UnitType:           item.GetUnitType(),
		})
	}

	return purchases, int(pbRes.Msg.GetTotalRecords()), nil

}
func (s *CustomerService) ListBranches(ctx context.Context, contactId string) ([]*domain.BranchSummary, error) {
	pbReq := customerv1.GetBranchesForContactRequest_builder{ContactId: proto.String(contactId)}.Build()

	pbRes, err := s.Client.GetBranchesForContact(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	var branches []*domain.BranchSummary

	for _, branch := range pbRes.Msg.GetBranches() {
		branches = append(branches, &domain.BranchSummary{
			Id:   branch.GetId(),
			Name: branch.GetName(),
		})
	}

	return branches, nil
}
