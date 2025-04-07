package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
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
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListOrderPackingList(ctx context.Context, params oas.ListOrderPackingListParams) (*oas.ListOrderPackingListOK, error) {
	packingLists, err := h.service.Order.ListPackingListsByOrder(ctx, params.ID)

	if err != nil {
		return nil, err
	}
	res := &oas.ListOrderPackingListOK{}
	for _, packingList := range packingLists {
		res.PackingLists = append(res.PackingLists, oas.PackingListSummary{
			InvoiceID:    packingList.InvoiceId,
			DateInvoiced: packingList.DateInvoiced,
		})
	}
	return res, nil

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

func (h Handler) GetOrderInvoices(ctx context.Context, params oas.GetOrderInvoicesParams) (*oas.GetOrderInvoicesOK, error) {
	invoices, err := h.service.Billing.GetInvoicesByOrder(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	res := &oas.GetOrderInvoicesOK{}
	for _, invoice := range invoices {
		res.Invoices = append(res.Invoices, oas.InvoiceSummary{
			ID:           invoice.Id,
			OrderID:      invoice.OrderId,
			DateInvoiced: invoice.DateInvoiced,
			TotalAmount:  invoice.TotalAmount,
		})
	}
	return res, nil
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

func (h Handler) ListCustomerBranches(ctx context.Context, params oas.ListCustomerBranchesParams) (oas.ListCustomerBranchesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListOrders(ctx context.Context, params oas.ListOrdersParams) (oas.ListOrdersRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	pbReq := orderv1.ListOrdersRequest_builder{
		Page:     proto.Int32(int32(params.Page)),
		PageSize: proto.Int32(int32(params.PageSize)),
		BranchId: proto.String(userSession.Profile.BranchID),
	}.Build()

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
	pbReq := orderv1.GetOrderRequest_builder{Id: proto.String(params.ID)}.Build()
	pbRes, err := h.service.Order.Client.GetOrder(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	if pbRes.Msg.GetOrder().GetBranchId() != userSession.Profile.BranchID {
		return nil, fmt.Errorf("user is not authorized to access this branch")
	}

	packingLists, err := h.service.Order.ListPackingListsByOrder(ctx, params.ID)
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

	for _, packingList := range packingLists {
		response.Order.PackingLists = append(response.Order.PackingLists, oas.PackingListSummary{
			InvoiceID:    packingList.InvoiceId,
			DateInvoiced: packingList.DateInvoiced,
		})
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
