package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	orderv1 "github.com/materials-resources/customer-api/internal/grpc-client/order"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/service"
	"github.com/materials-resources/customer-api/internal/session"
	"github.com/materials-resources/customer-api/internal/zitadel"
)

func NewHandler(service service.Service, sessionManager *session.Manager) Handler {
	z, err := zitadel.NewZitadelClient()
	if err != nil {
		panic(err)
	}
	return Handler{
		sessionManager: sessionManager,
		z:              z,
		service:        service,
	}
}

type Handler struct {
	sessionManager *session.Manager
	service        service.Service
	z              *zitadel.Client
}

func (h Handler) GetQuote(ctx context.Context, params oas.GetQuoteParams) (*oas.GetQuoteOK, error) {
	userSession := h.sessionManager.GetUserSession(ctx)
	pbReq := &orderv1.GetQuoteRequest{Id: params.ID}
	pbRes, err := h.service.Order.Client.GetQuote(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	if pbRes.Msg.GetQuote().GetBranch().GetId() != userSession.Profile.BranchID {
		fmt.Println("user is not authorized to access this branch")
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
	res, err := h.service.Order.GetQuote(ctx, params)
	return res, err
}

func (h Handler) ListQuotes(ctx context.Context, params oas.ListQuotesParams) (*oas.ListQuotesOK, error) {
	userSession := h.sessionManager.GetUserSession(ctx)
	pbReq := &orderv1.ListQuotesRequest{
		Page:     int32(params.Page),
		PageSize: int32(params.PageSize),
		BranchId: userSession.Profile.BranchID,
	}

	pbRes, err := h.service.Order.Client.ListQuotes(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	response := oas.ListQuotesOK{
		TotalRecords: int(pbRes.Msg.GetTotalRecords()),
	}

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
	return &response, err
}

func (h Handler) ListOrders(ctx context.Context, params oas.ListOrdersParams) (*oas.ListOrdersOK, error) {
	userSession := h.sessionManager.GetUserSession(ctx)
	pbReq := &orderv1.ListOrdersRequest{
		Page:     int32(params.Page),
		PageSize: int32(params.PageSize),
		BranchId: userSession.Profile.BranchID,
	}

	pbRes, err := h.service.Order.Client.ListOrders(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	response := oas.ListOrdersOK{
		TotalRecords: int(pbRes.Msg.GetTotalRecords()),
	}

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

func (h Handler) SetActiveBranch(ctx context.Context, req *oas.SetActiveBranchReq) (oas.SetActiveBranchRes, error) {
	fmt.Println("changed branch")
	// check if user can access this branch
	// get user id from token

	// update the branch
	h.z.ChangeUserBranchId(ctx, "295379791043934934", req.GetBranchID())
	// return success

	return &oas.SetActiveBranchOK{}, nil
}

func (h Handler) CreateQuote(ctx context.Context, req *oas.CreateQuoteReq) (oas.CreateQuoteRes, error) {
	return h.service.Order.CreateQuote(ctx, req)
}

func (h Handler) GetOrder(ctx context.Context, params oas.GetOrderParams) (oas.GetOrderRes, error) {
	userSession := h.sessionManager.GetUserSession(ctx)
	pbReq := &orderv1.GetOrderRequest{Id: params.ID}
	pbRes, err := h.service.Order.Client.GetOrder(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	if pbRes.Msg.GetOrder().GetBranchId() != userSession.Profile.BranchID {
		return nil, fmt.Errorf("user is not authorized to access this branch")
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

func (h Handler) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	res, err := h.service.Catalog.GetProduct(ctx, params)
	return res, err
}

func (h Handler) ListCustomerBranches(ctx context.Context, params oas.ListCustomerBranchesParams) (*oas.ListCustomerBranchesOK, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (*oas.SearchProductsOK, error) {
	return h.service.Search.SearchProducts(ctx, req)
}

var _ oas.Handler = (*Handler)(nil)

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
