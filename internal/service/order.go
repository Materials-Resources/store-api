package service

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"context"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/internal/domain"
	"github.com/materials-resources/store-api/internal/oas"
	orderv1 "github.com/materials-resources/store-api/internal/proto/order"
	"github.com/materials-resources/store-api/internal/proto/order/orderconnect"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
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
	pbReq := orderv1.GetQuoteRequest_builder{Id: proto.String(req.ID)}.Build()
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

func (s *Order) GetOrder(ctx context.Context, orderId string) (*domain.Order, error) {
	pbReq := orderv1.GetOrderRequest_builder{Id: proto.String(orderId)}.Build()
	pbRes, err := s.Client.GetOrder(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	order := &domain.Order{
		Id:              pbRes.Msg.GetOrder().GetId(),
		ContactId:       pbRes.Msg.GetOrder().GetContactId(),
		BranchId:        pbRes.Msg.GetOrder().GetBranchId(),
		PurchaseOrder:   pbRes.Msg.GetOrder().GetPurchaseOrder(),
		Status:          convertOrderStatus(pbRes.Msg.GetOrder().GetStatus()),
		DateCreated:     pbRes.Msg.GetOrder().GetDateCreated().AsTime(),
		DateRequested:   pbRes.Msg.GetOrder().GetDateRequested().AsTime(),
		Taker:           pbRes.Msg.GetOrder().GetTaker(),
		ShippingAddress: domain.Address{},
	}

	if pbRes.Msg.GetOrder().GetShippingAddress() != nil {
		order.ShippingAddress = domain.Address{
			Name:       pbRes.Msg.GetOrder().GetShippingAddress().GetName(),
			LineOne:    pbRes.Msg.GetOrder().GetShippingAddress().GetLineOne(),
			LineTwo:    pbRes.Msg.GetOrder().GetShippingAddress().GetLineTwo(),
			City:       pbRes.Msg.GetOrder().GetShippingAddress().GetCity(),
			State:      pbRes.Msg.GetOrder().GetShippingAddress().GetState(),
			PostalCode: pbRes.Msg.GetOrder().GetShippingAddress().GetPostalCode(),
			Country:    pbRes.Msg.GetOrder().GetShippingAddress().GetCountry(),
		}
	}

	for _, itemPb := range pbRes.Msg.GetOrder().GetOrderItems() {

		item := &domain.OrderItem{

			ProductId:         itemPb.GetProductId(),
			ProductSn:         itemPb.GetProductSn(),
			ProductName:       itemPb.GetProductName(),
			CustomerProductSn: itemPb.GetCustomerProductSn(),
			OrderedQuantity:   itemPb.GetOrderedQuantity(),
			ShippedQuantity:   itemPb.GetShippedQuantity(),
			RemainingQuantity: itemPb.GetRemainingQuantity(),
			UnitType: domain.UnitOfMeasurement{
				Id: itemPb.GetUnitType(),
			},
			UnitPrice:  itemPb.GetUnitPrice(),
			TotalPrice: itemPb.GetTotalPrice(),
		}

		order.Items = append(order.Items, item)
	}

	return order, nil
}

func (s *Order) ListOrders(ctx context.Context, page, pageSize int32, branchId string) ([]*domain.OrderSummary, int32,
	error) {
	pbReq := orderv1.ListOrdersRequest_builder{
		Page:     proto.Int32(page),
		PageSize: proto.Int32(pageSize),
		BranchId: proto.String(branchId),
	}.Build()

	pbRes, err := s.Client.ListOrders(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, 0, err
	}

	var orders []*domain.OrderSummary

	for _, pbOrder := range pbRes.Msg.GetOrders() {
		orders = append(orders, &domain.OrderSummary{
			Id:            pbOrder.GetId(),
			ContactId:     pbOrder.GetContactId(),
			BranchId:      pbOrder.GetBranchId(),
			PurchaseOrder: pbOrder.GetPurchaseOrder(),
			Status:        convertOrderStatus(pbOrder.GetStatus()),
			DateCreated:   pbOrder.GetDateCreated().AsTime(),
			DateRequested: pbOrder.GetDateRequested().AsTime(),
		})
	}
	return orders, pbRes.Msg.GetTotalRecords(), nil
}

func (s *Order) ListQuotes(ctx context.Context, req oas.ListQuotesParams) (*oas.ListQuotesOK, error) {
	pbReq := orderv1.ListQuotesRequest_builder{
		Page:     proto.Int32(int32(req.Page)),
		PageSize: proto.Int32(int32(req.PageSize)),
		BranchId: proto.String("100039"),
	}.Build()

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

func (s *Order) ListPackingListsByOrder(ctx context.Context, orderId string) ([]*domain.PackingListSummary, error) {
	pbReq := orderv1.ListPackingListsByOrderRequest_builder{
		OrderId: proto.String(orderId),
	}.Build()

	pbRes, err := s.Client.ListPackingListsByOrder(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	var packingLists []*domain.PackingListSummary

	for _, pbPackingList := range pbRes.Msg.GetPackingLists() {
		packingLists = append(packingLists, &domain.PackingListSummary{
			InvoiceId:    pbPackingList.GetInvoiceId(),
			OrderId:      pbPackingList.GetOrderId(),
			DateInvoiced: pbPackingList.GetDateInvoiced().AsTime(),
		})
	}

	return packingLists, nil

}

func convertOrderStatus(status orderv1.OrderStatus) domain.OrderStatus {
	switch status {
	case orderv1.OrderStatus_ORDER_STATUS_COMPLETED:
		return domain.OrderStatusCompleted
	case orderv1.OrderStatus_ORDER_STATUS_PENDING_APPROVAL:
		return domain.OrderStatusPendingApproval
	case orderv1.OrderStatus_ORDER_STATUS_APPROVED:
		return domain.OrderStatusApproved
	case orderv1.OrderStatus_ORDER_STATUS_CANCELLED:
		return domain.OrderStatusCancelled
	case orderv1.OrderStatus_ORDER_STATUS_UNSPECIFIED:
		return domain.OrderStatusUnspecified
	default:
		return domain.OrderStatusUnspecified
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
