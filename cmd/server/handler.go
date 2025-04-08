package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/materials-resources/store-api/internal/domain"
	"github.com/materials-resources/store-api/internal/mailer"
	"github.com/materials-resources/store-api/internal/oas"
	customerv1 "github.com/materials-resources/store-api/internal/proto/customer"
	orderv1 "github.com/materials-resources/store-api/internal/proto/order"
	"github.com/materials-resources/store-api/internal/service"
	"github.com/materials-resources/store-api/internal/session"
	"github.com/materials-resources/store-api/internal/zitadel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ oas.Handler = (*Handler)(nil)

func NewHandler(service service.Service, sessionManager *session.Manager) Handler {
	z, err := zitadel.NewZitadelClient()
	if err != nil {
		panic(err)
	}
	m := mailer.New("smtp.materialsresources.or", 25, "", "", "noreply@materialsresources.org")
	return Handler{
		sessionManager: sessionManager,
		z:              z,
		service:        service,
		mailer:         m,
	}
}

type Handler struct {
	sessionManager *session.Manager
	service        service.Service
	z              *zitadel.Client
	mailer         mailer.Mailer
}

func (h Handler) GetPackingListReport(ctx context.Context, params oas.GetPackingListReportParams) (oas.GetPackingListReportRes, error) {
	data, err := h.service.Report.GetPackingList(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return &oas.GetPackingListReportOK{Data: data}, nil
}

func (h Handler) ContactUs(ctx context.Context, req *oas.ContactUsReq) error {
	d := map[string]any{
		"Organization": req.GetOrganization(),
		"Name":         req.GetName(),
		"Email":        req.GetEmail(),
		"Message":      req.GetMessage(),
	}
	err := h.mailer.Send("smallegan@emrsinc.com", "smallegan@emrsinc.com", "contact_request.tmpl", d)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (h Handler) GetInvoiceReport(ctx context.Context, params oas.GetInvoiceReportParams) (oas.GetInvoiceReportRes,
	error) {
	data, err := h.service.Report.GetInvoice(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return &oas.GetInvoiceReportOK{
		Data: data,
	}, nil
}

func (h Handler) GetActiveBranches(ctx context.Context) (oas.GetActiveBranchesRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	pbReq := customerv1.GetBranchRequest_builder{
		Id: proto.String(userSession.Profile.BranchID),
	}.Build()

	pbRes, err := h.service.Customer.Client.GetBranch(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}
	response := oas.GetActiveBranchesOK{
		Branch: oas.Branch{
			ID:   pbRes.Msg.GetBranch().GetId(),
			Name: pbRes.Msg.GetBranch().GetName(),
		},
	}
	return &response, nil
}

func (h Handler) GetQuote(ctx context.Context, params oas.GetQuoteParams) (oas.GetQuoteRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	pbReq := orderv1.GetQuoteRequest_builder{Id: proto.String(params.ID)}.Build()
	pbRes, err := h.service.Order.Client.GetQuote(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	if pbRes.Msg.GetQuote().GetBranch().GetId() != userSession.Profile.BranchID {
		return nil, fmt.Errorf("user is not authorized to access this branch")
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

func (h Handler) GetRecentPurchases(ctx context.Context) (*oas.GetRecentPurchasesOK, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	pbReq := customerv1.GetRecentPurchasesByBranchRequest_builder{
		Id:    proto.String(userSession.Profile.BranchID),
		Limit: proto.Int32(10),
	}.Build()
	pbRes, err := h.service.Customer.Client.GetRecentPurchasesByBranch(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}
	response := oas.GetRecentPurchasesOK{}

	for _, purchase := range pbRes.Msg.GetItems() {
		response.Purchases = append(response.Purchases, oas.PurchaseSummary{
			ProductID:          purchase.GetProductId(),
			ProductSn:          purchase.GetProductSn(),
			ProductName:        purchase.GetProductName(),
			ProductDescription: purchase.GetProductDescription(),
			OrderedQuantity:    purchase.GetOrderedQuantity(),
			UnitOfMeasurement:  purchase.GetUnitType(),
		})
	}

	return &response, nil

}

func (h Handler) ListCustomerBranches(ctx context.Context) (oas.ListCustomerBranchesRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	branches, err := h.service.Customer.ListBranches(ctx, userSession.Profile.ContactID)
	if err != nil {
		return nil, err
	}
	response := oas.ListCustomerBranchesOK{}
	for _, branch := range branches {
		response.Branches = append(response.Branches, oas.Branch{
			ID:   branch.Id,
			Name: branch.Name,
		})
	}
	return &response, nil
}

func (h Handler) ListOrders(ctx context.Context, params oas.ListOrdersParams) (oas.ListOrdersRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	orders, total, err := h.service.Order.ListOrders(ctx, int32(params.Page), int32(params.PageSize),
		userSession.Profile.BranchID)

	if err != nil {
		return nil, err
	}

	response := oas.ListOrdersOK{
		TotalRecords: int(total),
	}

	for _, order := range orders {
		response.Orders = append(response.Orders, oas.OrderSummary{
			ID:            order.Id,
			BranchID:      order.BranchId,
			ContactID:     order.ContactId,
			PurchaseOrder: order.PurchaseOrder,
			Status:        mapOrderStatus(order.Status),
			DateCreated:   order.DateCreated,
			DateRequested: order.DateRequested,
		})
	}
	return &response, nil
}

func (h Handler) ListQuotes(ctx context.Context, params oas.ListQuotesParams) (oas.ListQuotesRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	pbReq := orderv1.ListQuotesRequest_builder{
		Page:     proto.Int32(int32(params.Page)),
		PageSize: proto.Int32(int32(params.PageSize)),
		BranchId: proto.String(userSession.Profile.BranchID),
	}.Build()

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

func (h Handler) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (oas.SearchProductsRes, error) {
	return h.service.Search.SearchProducts(ctx, req)
}

func (h Handler) SetActiveBranch(ctx context.Context, req *oas.SetActiveBranchReq) (oas.SetActiveBranchRes, error) {
	// check if user can access this branch
	// get user id from token

	// update the branch
	err := h.z.ChangeUserBranchId(ctx, "295379791043934934", req.GetBranchID())
	if err != nil {
		return nil, err
	}
	// return success

	return &oas.SetActiveBranchOK{}, nil
}

func (h Handler) CreateQuote(ctx context.Context, req *oas.CreateQuoteReq) (oas.CreateQuoteRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	pbReq := orderv1.CreateQuoteRequest_builder{
		BranchId:      proto.String(userSession.Profile.BranchID),
		ContactId:     proto.String(userSession.Profile.ContactID),
		Notes:         proto.String(req.Notes),
		RequestedDate: timestamppb.New(req.GetDateRequested()),
	}

	for _, item := range req.Items {
		pbReq.Items = append(pbReq.Items, orderv1.CreateQuoteRequest_Item_builder{
			ProductId: proto.String(item.GetProductID()),
			Quantity:  proto.Float64(item.GetQuantity()),
		}.Build())
	}

	pbRes, err := h.service.Order.Client.CreateQuote(ctx, connect.NewRequest(pbReq.Build()))
	if err != nil {
		return nil, err
	}

	return &oas.CreateQuoteCreated{
		QuoteID: pbRes.Msg.GetId(),
	}, nil
}

func (h Handler) GetOrder(ctx context.Context, params oas.GetOrderParams) (oas.GetOrderRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	order, err := h.service.Order.GetOrder(ctx, params.ID)

	if err != nil {
		return nil, err
	}

	if order.BranchId != userSession.Profile.BranchID {
		return nil, fmt.Errorf("user is not authorized to access this branch")
	}

	packingLists, err := h.service.Order.ListPackingListsByOrder(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	invoices, err := h.service.Billing.GetInvoicesByOrder(ctx, params.ID)

	oapiOrder := oas.Order{
		ID:            order.Id,
		ContactID:     order.ContactId,
		BranchID:      order.BranchId,
		PurchaseOrder: order.PurchaseOrder,
		Status:        mapOrderStatus(order.Status),
		DateCreated:   order.DateCreated,
		DateRequested: order.DateRequested,
		Taker:         oas.NewOptString(order.Taker),
		ShippingAddress: oas.Address{
			Name:       order.ShippingAddress.Name,
			LineOne:    order.ShippingAddress.LineOne,
			LineTwo:    order.ShippingAddress.LineTwo,
			City:       order.ShippingAddress.City,
			State:      order.ShippingAddress.State,
			PostalCode: order.ShippingAddress.PostalCode,
			Country:    order.ShippingAddress.Country,
		},
	}

	for _, packingList := range packingLists {
		oapiOrder.PackingLists = append(oapiOrder.PackingLists, oas.PackingListSummary{
			InvoiceID:    packingList.InvoiceId,
			DateInvoiced: packingList.DateInvoiced,
		})
	}

	for _, invoice := range invoices {
		oapiOrder.Invoices = append(oapiOrder.Invoices, oas.InvoiceSummary{
			ID:           invoice.Id,
			DateInvoiced: invoice.DateInvoiced,
			TotalAmount:  invoice.TotalAmount,
			PaidAmount:   invoice.PaidAmount,
		})
	}

	for _, item := range order.Items {
		oapiOrder.Items = append(oapiOrder.Items, oas.OrderItem{
			ProductID:         item.ProductId,
			ProductSn:         item.ProductSn,
			ProductName:       item.ProductName,
			CustomerProductSn: item.CustomerProductSn,
			OrderedQuantity:   item.OrderedQuantity,
			ShippedQuantity:   item.ShippedQuantity,
			RemainingQuantity: item.RemainingQuantity,
			UnitType:          item.UnitType.Id,
			UnitPrice:         item.UnitPrice,
			TotalPrice:        item.TotalPrice,
		})
	}

	response := oas.GetOrderOK{
		Order: oapiOrder,
	}

	return &response, nil
}

func (h Handler) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	product, err := h.service.Catalog.GetProduct(ctx, params)

	if err != nil {
		return nil, err
	}

	res := oas.GetProductOK{
		Product: oas.Product{
			ID:               product.Id,
			Sn:               product.Sn,
			Name:             product.Name,
			Description:      product.Description,
			ProductGroupName: product.ProductGroupName,
			ProductGroupID:   product.ProductGroupId,
			SalesUnitOfMeasurement: oas.UnitOfMeasurement{
				ID:               product.SalesUnitOfMeasurement.Id,
				ConversionFactor: product.SalesUnitOfMeasurement.ConversionFactor,
			},
		},
	}

	return &res, err
}

func mapOrderStatus(status domain.OrderStatus) oas.OrderStatus {
	switch status {
	case domain.OrderStatusCompleted:
		return oas.OrderStatusCompleted
	case domain.OrderStatusPendingApproval:
		return oas.OrderStatusPendingApproval
	case domain.OrderStatusApproved:
		return oas.OrderStatusApproved
	case domain.OrderStatusCancelled:
		return oas.OrderStatusCancelled
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
