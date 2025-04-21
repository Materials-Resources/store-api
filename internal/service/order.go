package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/internal/domain"
	orderv1 "github.com/materials-resources/store-api/internal/proto/order"
	"github.com/materials-resources/store-api/internal/proto/order/orderconnect"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type Order struct {
	Client orderconnect.OrderServiceClient
}

func NewOrderService(a *app.App) *Order {
	otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
	if err != nil {
		a.Logger.Fatal().Str("service", "order").Err(err).Msg("could not create otel interceptor")
	}

	return &Order{
		Client: orderconnect.NewOrderServiceClient(http.DefaultClient,
			a.Config.Services.OrderUrl,
			connect.WithInterceptors(otelInterceptor),
			connect.WithGRPC()),
	}
}

func (s *Order) GetQuote(ctx context.Context, id string) (*domain.Quote, error) {
	pbReq := orderv1.GetQuoteRequest_builder{Id: proto.String(id)}.Build()
	pbRes, err := s.Client.GetQuote(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	quote := &domain.Quote{
		Id:            pbRes.Msg.GetQuote().GetId(),
		BranchId:      pbRes.Msg.GetQuote().GetBranch().GetId(),
		PurchaseOrder: pbRes.Msg.GetQuote().GetPurchaseOrder(),
		Status:        convertQuoteStatus(pbRes.Msg.GetQuote().GetStatus()),
		DateCreated:   pbRes.Msg.GetQuote().GetDateCreated().AsTime(),
	}

	for _, itemPb := range pbRes.Msg.GetQuote().GetItems() {
		quote.Items = append(quote.Items, &domain.QuoteItem{
			ProductId:         itemPb.GetProductId(),
			ProductSn:         itemPb.GetProductSn(),
			ProductName:       itemPb.GetProductName(),
			CustomerProductSn: itemPb.GetCustomerProductSn(),
			UnitPrice:         itemPb.GetUnitPrice(),
			UnitType: domain.UnitOfMeasurement{
				Id: itemPb.GetUnitType(),
			},
			OrderedQuantity: itemPb.GetOrderedQuantity(),
			TotalPrice:      itemPb.GetTotalPrice(),
		})
	}

	return quote, nil
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
		DateOrdered:     pbRes.Msg.GetOrder().GetDateOrdered().AsTime(),
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
			DateOrdered:   pbOrder.GetDateOrdered().AsTime(),
			DateRequested: pbOrder.GetDateRequested().AsTime(),
		})
	}
	return orders, pbRes.Msg.GetTotalRecords(), nil
}

func (s *Order) ListQuotes(ctx context.Context, page, pageSize int32, branchId string) ([]*domain.QuoteSummary, int32, error) {
	pbReq := orderv1.ListQuotesRequest_builder{
		Page:     proto.Int32(page),
		PageSize: proto.Int32(pageSize),
		BranchId: proto.String(branchId),
	}.Build()

	pbRes, err := s.Client.ListQuotes(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, 0, err
	}

	var quotes []*domain.QuoteSummary
	for _, pbQuote := range pbRes.Msg.GetQuotes() {
		quotes = append(quotes, &domain.QuoteSummary{
			Id:            pbQuote.GetId(),
			BranchId:      pbQuote.GetBranch().GetId(),
			ContactId:     pbQuote.GetContact().GetId(),
			PurchaseOrder: pbQuote.GetPurchaseOrder(),
			Status:        convertQuoteStatus(pbQuote.GetStatus()),
			DateCreated:   pbQuote.GetDateCreated().AsTime(),
			DateExpires:   pbQuote.GetDateExpires().AsTime(),
		})
	}

	return quotes, pbRes.Msg.GetTotalRecords(), nil
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

func convertQuoteStatus(status orderv1.QuoteStatus) domain.QuoteStatus {
	switch status {
	case orderv1.QuoteStatus_QUOTE_STATUS_PROCESSED:
		return domain.QuoteStatusProcessed
	case orderv1.QuoteStatus_QUOTE_STATUS_CANCELLED:
		return domain.QuoteStatusCancelled
	case orderv1.QuoteStatus_QUOTE_STATUS_PENDING_APPROVAL:
		return domain.QuoteStatusPending
	case orderv1.QuoteStatus_QUOTE_STATUS_CLOSED:
		return domain.QuoteStatusClosed
	case orderv1.QuoteStatus_QUOTE_STATUS_UNSPECIFIED:
		return domain.QuoteStatusUnspecified
	default:
		return domain.QuoteStatusUnspecified
	}
}
