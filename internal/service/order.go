package service

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"context"
	"github.com/materials-resources/store-api/app"
	orderv1 "github.com/materials-resources/store-api/internal/grpc-client/order"
	"github.com/materials-resources/store-api/internal/grpc-client/order/orderconnect"
	"github.com/materials-resources/store-api/internal/oas"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
)

type Order struct {
	Client orderconnect.OrderServiceClient
}

func newInterceptor(tp trace.TracerProvider, mp metric.MeterProvider, p propagation.TextMapPropagator) (connect.Interceptor, error) {
	return otelconnect.NewInterceptor(
		otelconnect.WithTracerProvider(tp),
		otelconnect.WithMeterProvider(mp),
		otelconnect.WithPropagator(p),
	)
}

func NewOrderService(a *app.App) *Order {
	otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
	if err != nil {
		log.Fatal(err)
	}

	return &Order{
		Client: orderconnect.NewOrderServiceClient(http.DefaultClient,
			"http://localhost:8082",
			connect.WithInterceptors(otelInterceptor),
			connect.WithGRPC()),
	}
}

func (s *Order) GetQuote(ctx context.Context, req oas.GetQuoteParams) (*oas.GetQuoteOK, error) {
	pbReq := &orderv1.GetQuoteRequest{Id: req.ID}
	pbRes, err := s.Client.GetQuote(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	response := oas.GetQuoteOK{
		Quote: oas.Quote{
			ID:            pbRes.Msg.GetQuote().GetId(),
			PurchaseOrder: pbRes.Msg.GetQuote().GetPurchaseOrder(),
			Status:        convertQuoteStatus(pbRes.Msg.GetQuote().GetStatus()),
			DateCreated:   pbRes.Msg.GetQuote().GetDateCreated().AsTime(),
		},
	}

	for _, item := range pbRes.Msg.GetQuote().GetItems() {
		response.Quote.Items = append(response.Quote.Items, oas.QuoteItem{
			ProductID:         item.GetProductId(),
			ProductSn:         item.GetProductSn(),
			ProductName:       item.GetProductName(),
			CustomerProductSn: item.GetCustomerProductSn(),
			UnitPrice:         item.GetUnitPrice(),
			UnitType:          item.GetUnitType(),
			OrderedQuantity:   item.GetOrderedQuantity(),
			TotalPrice:        item.GetTotalPrice(),
		})
	}

	return &response, nil
}

func (s *Order) GetOrder(ctx context.Context, req oas.GetOrderParams) (oas.GetOrderRes, error) {
	pbReq := &orderv1.GetOrderRequest{Id: req.ID}
	pbRes, err := s.Client.GetOrder(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	response := oas.GetOrderOK{
		Order: oas.Order{
			ID:              pbRes.Msg.GetOrder().GetId(),
			ContactID:       pbRes.Msg.GetOrder().GetContactId(),
			BranchID:        pbRes.Msg.GetOrder().GetBranchId(),
			PurchaseOrder:   pbRes.Msg.GetOrder().GetPurchaseOrder(),
			Status:          convertOrderStatus(pbRes.Msg.GetOrder().GetStatus()),
			DateCreated:     pbRes.Msg.GetOrder().GetDateCreated().AsTime(),
			DateRequested:   pbRes.Msg.GetOrder().GetDateRequested().AsTime(),
			Taker:           oas.NewOptString(""),
			ShippingAddress: oas.Address{},
			Total:           0,
		},
	}

	response.Order.SetShippingAddress(oas.Address{
		ID:         "",
		Name:       pbRes.Msg.GetOrder().GetShippingAddress().GetName(),
		LineOne:    pbRes.Msg.GetOrder().GetShippingAddress().GetLineOne(),
		LineTwo:    pbRes.Msg.GetOrder().GetShippingAddress().GetLineTwo(),
		City:       pbRes.Msg.GetOrder().GetShippingAddress().GetCity(),
		State:      pbRes.Msg.GetOrder().GetShippingAddress().GetState(),
		PostalCode: pbRes.Msg.GetOrder().GetShippingAddress().GetPostalCode(),
		Country:    pbRes.Msg.GetOrder().GetShippingAddress().GetCountry(),
	})

	for _, item := range pbRes.Msg.GetOrder().GetOrderItems() {
		response.Order.Items = append(response.Order.Items, oas.OrderItem{
			ProductSn:           item.GetProductSn(),
			ProductName:         item.GetProductName(),
			ProductID:           item.GetProductId(),
			CustomerProductSn:   item.GetCustomerProductSn(),
			OrderedQuantity:     item.GetOrderedQuantity(),
			ShippedQuantity:     item.GetShippedQuantity(),
			BackOrderedQuantity: item.GetBackOrderedQuantity(),
			UnitType:            item.GetUnitType(),
			UnitPrice:           item.GetUnitPrice(),
			TotalPrice:          item.GetTotalPrice(),
		})
	}

	return &response, nil
}

func (s *Order) ListOrders(ctx context.Context, req oas.ListOrdersParams) (*oas.ListOrdersOK, error) {
	pbReq := &orderv1.ListOrdersRequest{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		BranchId: "100039",
	}

	pbRes, err := s.Client.ListOrders(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	response := oas.ListOrdersOK{}

	for _, pbOrder := range pbRes.Msg.GetOrders() {
		response.Orders = append(response.Orders, oas.OrderSummary{
			ID:            pbOrder.GetId(),
			ContactID:     pbOrder.GetContactId(),
			BranchID:      pbOrder.GetBranchId(),
			PurchaseOrder: pbOrder.GetPurchaseOrder(),
			Status:        convertOrderStatus(pbOrder.GetStatus()),
			DateCreated:   pbOrder.GetDateCreated().AsTime(),
			DateRequested: pbOrder.GetDateRequested().AsTime(),
		})
	}
	return &response, nil
}

func (s *Order) ListQuotes(ctx context.Context, req oas.ListQuotesParams) (*oas.ListQuotesOK, error) {
	pbReq := &orderv1.ListQuotesRequest{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		BranchId: "100039",
	}

	pbRes, err := s.Client.ListQuotes(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	response := oas.ListQuotesOK{}

	for _, pbQuote := range pbRes.Msg.GetQuotes() {
		response.Quotes = append(response.Quotes, oas.QuoteSummary{
			ID:            pbQuote.GetId(),
			BranchID:      pbQuote.GetBranch().GetId(),
			ContactID:     pbQuote.GetContact().GetId(),
			PurchaseOrder: pbQuote.GetPurchaseOrder(),
			Status:        convertQuoteStatus(pbQuote.GetStatus()),
			DateCreated:   pbQuote.GetDateCreated().AsTime(),
			DateExpires:   pbQuote.GetDateExpires().AsTime(),
		})
	}

	return &response, nil
}

func convertOrderStatus(status orderv1.OrderStatus) oas.OrderStatus {
	switch status {
	case orderv1.OrderStatus_ORDER_STATUS_COMPLETED:
		return oas.OrderStatusCompleted
	case orderv1.OrderStatus_ORDER_STATUS_PENDING_APPROVAL:
		return oas.OrderStatusPendingApproval
	case orderv1.OrderStatus_ORDER_STATUS_APPROVED:
		return oas.OrderStatusApproved
	case orderv1.OrderStatus_ORDER_STATUS_CANCELLED:
		return oas.OrderStatusCancelled
	case orderv1.OrderStatus_ORDER_STATUS_UNSPECIFIED:
		return oas.OrderStatusUnspecified
	default:
		return oas.OrderStatusUnspecified
	}
}

func convertQuoteStatus(status orderv1.QuoteStatus) oas.QuoteStatus {
	switch status {
	case orderv1.QuoteStatus_QUOTE_STATUS_APPROVED:
		return oas.QuoteStatusApproved
	case orderv1.QuoteStatus_QUOTE_STATUS_CANCELLED:
		return oas.QuoteStatusCancelled
	case orderv1.QuoteStatus_QUOTE_STATUS_PENDING_APPROVAL:
		return oas.QuoteStatusPendingApproval
	case orderv1.QuoteStatus_QUOTE_STATUS_EXPIRED:
		return oas.QuoteStatusExpired
	case orderv1.QuoteStatus_QUOTE_STATUS_UNSPECIFIED:
		return oas.QuoteStatusUnspecified
	default:
		return oas.QuoteStatusUnspecified
	}
}
