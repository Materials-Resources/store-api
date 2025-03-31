package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/internal/domain"
	billingv1 "github.com/materials-resources/store-api/internal/proto/billing"
	"github.com/materials-resources/store-api/internal/proto/billing/billingconnect"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type BillingService struct {
	client billingconnect.BillingServiceClient
}

func NewBillingService() *BillingService {
	return &BillingService{
		client: billingconnect.NewBillingServiceClient(http.DefaultClient,
			"http://localhost:8082", connect.WithGRPC()),
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
		})

	}
	return invoices, nil
}
