package service

import (
	"connectrpc.com/connect"
	"context"
	orderv1 "github.com/materials-resources/customer-api/internal/grpc-client/order"
	"github.com/materials-resources/customer-api/internal/grpc-client/order/orderconnect"
	"github.com/materials-resources/customer-api/internal/oas"
	"net/http"
)

type Order struct {
	Client orderconnect.OrderServiceClient
}

func NewOrderService() *Order {
	return &Order{
		Client: orderconnect.NewOrderServiceClient(http.DefaultClient,
			"http://localhost:8082",
			connect.WithGRPC()),
	}
}

func (s *Order) GetOrder(ctx context.Context, req oas.GetOrderParams) (oas.GetOrderRes, error) {
	pbReq := &orderv1.GetOrderRequest{Id: req.ID}
	pbRes, err := s.Client.GetOrder(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	response := oas.GetOrderOK{
		Order: oas.Order{
			ID:                   pbRes.Msg.GetOrder().GetId(),
			ContactID:            pbRes.Msg.GetOrder().GetContactId(),
			BranchID:             pbRes.Msg.GetOrder().GetBranchId(),
			PurchaseOrder:        pbRes.Msg.GetOrder().GetPurchaseOrder(),
			Status:               convertOrderStatus(pbRes.Msg.GetOrder().GetStatus()),
			DateCreated:          pbRes.Msg.GetOrder().GetDateCreated().AsTime(),
			DateRequested:        pbRes.Msg.GetOrder().GetDateRequested().AsTime(),
			Taker:                oas.NewOptString(""),
			DeliveryInstructions: "",
			ShippingAddress:      oas.Address{},
			Total:                0,
		},
	}

	for _, item := range pbRes.Msg.GetOrder().GetOrderItems() {
		response.Order.Items = append(response.Order.Items, oas.OrderItem{
			ProductSn:           item.GetProductSn(),
			ProductName:         item.GetProductName(),
			ProductID:           item.GetProductId(),
			CustomerProductSn:   item.GetCustomerProductSn(),
			OrderQuantity:       0,
			OrderQuantityUnit:   item.GetQuantityUnit(),
			PriceUnit:           "",
			Price:               0,
			TotalPrice:          0,
			ShippedQuantity:     0,
			BackOrderedQuantity: 0,
		})
	}

	return &response, nil
}

func (s *Order) ListOrders(ctx context.Context, req oas.ListOrdersParams) (*oas.ListOrdersOK, error) {
	pbReq := &orderv1.ListOrdersRequest{
		Page:     1,
		PageSize: 50,
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

func convertOrderStatus(status orderv1.OrderStatus) oas.OrderStatus {
	switch status {
	case orderv1.OrderStatus_STATUS_COMPLETED:
		return oas.OrderStatusCompleted
	case orderv1.OrderStatus_STATUS_PENDING_APPROVAL:
		return oas.OrderStatusPendingApproval
	case orderv1.OrderStatus_STATUS_APPROVED:
		return oas.OrderStatusApproved
	case orderv1.OrderStatus_STATUS_CANCELLED:
		return oas.OrderStatusCancelled
	default:
		return oas.OrderStatusUnspecified
	}
}
