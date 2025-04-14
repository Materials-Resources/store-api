package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/internal/domain"
	billingv1 "github.com/materials-resources/store-api/internal/proto/billing"
	"github.com/materials-resources/store-api/internal/proto/billing/billingconnect"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type BillingService struct {
	client billingconnect.BillingServiceClient
}

func NewBillingService(a *app.App) *BillingService {
	otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
	if err != nil {
		a.Logger.Fatal().Str("service", "billing").Err(err).Msg("could not create otel interceptor")
	}
	return &BillingService{
		client: billingconnect.NewBillingServiceClient(http.DefaultClient,
			a.Config.Services.BillingUrl, connect.WithInterceptors(otelInterceptor), connect.WithGRPC()),
	}
}

func (s *BillingService) GetInvoicesByOrder(ctx context.Context, orderId string) ([]*domain.InvoiceSummary, error) {
	req := billingv1.GetInvoicesByOrderRequest_builder{OrderId: proto.String(orderId)}

	res, err := s.client.GetInvoicesByOrder(ctx, connect.NewRequest(req.Build()))
	if err != nil {
		return nil, err
	}

	var invoices []*domain.InvoiceSummary

	for _, invoice := range res.Msg.GetInvoices() {
		invoices = append(invoices, &domain.InvoiceSummary{
			Id:           invoice.GetId(),
			OrderId:      invoice.GetOrderId(),
			DateInvoiced: invoice.GetDateInvoiced().AsTime(),
			TotalAmount:  invoice.GetTotalAmount(),
			PaidAmount:   invoice.GetPaidAmount(),
		})

	}
	return invoices, nil
}

func (s *BillingService) GetInvoicesByBranch(ctx context.Context, branchId string, page,
	pageSize int32) ([]*domain.InvoiceSummary,
	error) {
	pbReq := billingv1.GetInvoicesByBranchRequest_builder{
		BranchId: proto.String(branchId), Page: proto.Int32(page),
		PageSize: proto.Int32(pageSize),
	}.Build()

	pbRes, err := s.client.GetInvoicesByBranch(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}
	var invoices []*domain.InvoiceSummary
	for _, invoice := range pbRes.Msg.GetInvoices() {
		invoices = append(invoices, &domain.InvoiceSummary{
			Id:           invoice.GetId(),
			OrderId:      invoice.GetOrderId(),
			DateInvoiced: invoice.GetDateInvoiced().AsTime(),
			TotalAmount:  invoice.GetTotalAmount(),
			PaidAmount:   invoice.GetPaidAmount(),
		})
	}
	return invoices, nil
}
