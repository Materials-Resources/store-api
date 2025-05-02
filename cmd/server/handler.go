package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/materials-resources/store-api/app"
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

func NewHandler(a *app.App, service service.Service, sessionManager *session.Manager, m mailer.Mailer) Handler {
	z, err := zitadel.NewZitadelClient(a)
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("could not create zitadel client")
	}
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

func (h Handler) ListInvoices(ctx context.Context, params oas.ListInvoicesParams) (oas.ListInvoicesRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	invoices, totalRecords, err := h.service.Billing.GetInvoicesByBranch(ctx, userSession.Profile.BranchID,
		int32(params.Page.Or(1)),
		int32(params.PageSize.Or(10)))

	if err != nil {
		return nil, err
	}

	response := &oas.ListInvoicesOK{
		TotalRecords: totalRecords,
		Invoices:     make([]oas.InvoiceSummary, 0),
	}

	for _, invoice := range invoices {
		response.Invoices = append(response.Invoices, oas.InvoiceSummary{
			ID:             invoice.Id,
			OrderID:        invoice.OrderId,
			DateInvoiced:   invoice.DateInvoiced,
			TotalAmount:    invoice.TotalAmount,
			PaidAmount:     invoice.PaidAmount,
			AdjustmentType: mapInvoiceAdjustmentType(invoice.AdjustmentType),
		})
	}
	return response, nil

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
		"Telephone":    req.GetTelephone(),
		"Message":      req.GetMessage(),
	}
	err := h.mailer.Send("contact_request.tmpl", d)
	if err != nil {
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

func (h Handler) GetActiveBranch(ctx context.Context) (oas.GetActiveBranchRes, error) {
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
	response := oas.GetActiveBranchOK{
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
	quote, err := h.service.Order.GetQuote(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if quote.BranchId != userSession.Profile.BranchID {
		return nil, fmt.Errorf("user is not authorized to access this branch")
	}

	oapiQuote := oas.Quote{
		ID:            quote.Id,
		PurchaseOrder: quote.PurchaseOrder,
		Status:        mapQuoteStatus(quote.Status),
		DateCreated:   quote.DateCreated,
	}

	for _, item := range quote.Items {
		oapiQuote.Items = append(oapiQuote.Items, oas.QuoteItem{
			ProductID:         item.ProductId,
			ProductSn:         item.ProductSn,
			ProductName:       item.ProductName,
			CustomerProductSn: item.CustomerProductSn,
			UnitPrice:         item.UnitPrice,
			UnitType:          item.UnitType.Id,
			OrderedQuantity:   item.OrderedQuantity,
			TotalPrice:        item.TotalPrice,
		})
	}

	response := oas.GetQuoteOK{
		Quote: oapiQuote,
	}
	return &response, err
}

func (h Handler) GetRecentPurchases(ctx context.Context, params oas.GetRecentPurchasesParams) (oas.GetRecentPurchasesRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	purchases, totalRecords, err := h.service.Customer.GetRecentPurchases(ctx, params.Page, params.PageSize, userSession.Profile.BranchID)

	if err != nil {
		return nil, err
	}

	response := oas.GetRecentPurchasesOK{
		TotalRecords: totalRecords,
		Purchases:    make([]oas.PurchaseSummary, 0),
	}

	for _, purchase := range purchases {
		response.Purchases = append(response.Purchases, oas.PurchaseSummary{
			ProductID:          purchase.ProductId,
			ProductSn:          purchase.ProductSn,
			ProductName:        purchase.ProductName,
			ProductDescription: purchase.ProductDescription,
			OrderedQuantity:    purchase.OrderedQuantity,
			UnitOfMeasurement:  purchase.UnitType,
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

	filters := &domain.OrderFilters{
		Id:            params.ID.Or(""),
		PurchaseOrder: params.PurchaseOrder.Or(""),
	}

	orders, total, err := h.service.Order.ListOrders(ctx, int32(params.Page), int32(params.PageSize),
		userSession.Profile.BranchID, filters)

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
			DateOrdered:   order.DateOrdered,
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

	filters := &domain.QuoteFilters{
		Id:            params.ID.Or(""),
		PurchaseOrder: params.PurchaseOrder.Or(""),
	}

	quotes, total, err := h.service.Order.ListQuotes(ctx, int32(params.Page), int32(params.PageSize),
		userSession.Profile.BranchID, filters)
	if err != nil {
		return nil, err
	}

	response := oas.ListQuotesOK{
		TotalRecords: int(total),
	}

	for _, quote := range quotes {
		response.Quotes = append(response.Quotes, oas.QuoteSummary{
			ID:            quote.Id,
			BranchID:      quote.BranchId,
			ContactID:     quote.ContactId,
			PurchaseOrder: quote.PurchaseOrder,
			Status:        mapQuoteStatus(quote.Status),
			DateCreated:   quote.DateCreated,
			DateExpires:   quote.DateExpires,
		})
	}

	return &response, err
}

func (h Handler) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (*oas.SearchProductsOK, error) {
	return h.service.Search.SearchProducts(ctx, req)
}

func (h Handler) SetActiveBranch(ctx context.Context, req *oas.SetActiveBranchReq) (oas.SetActiveBranchRes, error) {
	userSession, err := h.sessionManager.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	// update the branch
	err = h.z.ChangeUserBranchId(ctx, userSession.Profile.UserID, req.GetBranchID())
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
	if err != nil {
		return nil, err
	}

	oapiOrder := oas.Order{
		ID:            order.Id,
		ContactID:     order.ContactId,
		BranchID:      order.BranchId,
		PurchaseOrder: order.PurchaseOrder,
		Status:        mapOrderStatus(order.Status),
		DateOrdered:   order.DateOrdered,
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
		PackingLists: make([]oas.PackingListSummary, 0),
		Invoices:     make([]oas.InvoiceSummary, 0),
	}

	for _, packingList := range packingLists {
		oapiOrder.PackingLists = append(oapiOrder.PackingLists, oas.PackingListSummary{
			InvoiceID:    packingList.InvoiceId,
			DateInvoiced: packingList.DateInvoiced,
		})
	}

	for _, invoice := range invoices {
		oapiOrder.Invoices = append(oapiOrder.Invoices, oas.InvoiceSummary{
			ID:             invoice.Id,
			DateInvoiced:   invoice.DateInvoiced,
			TotalAmount:    invoice.TotalAmount,
			PaidAmount:     invoice.PaidAmount,
			AdjustmentType: mapInvoiceAdjustmentType(invoice.AdjustmentType),
		})
	}

	for _, item := range order.Items {
		oapiItem := oas.OrderItem{
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
			Disposition:       oas.NewOptOrderItemDisposition(mapOrderItemDisposition(item.Disposition)),
			Releases:          make([]oas.OrderItemRelease, 0),
		}

		for _, release := range item.Releases {
			oapiItem.Releases = append(oapiItem.Releases, oas.OrderItemRelease{
				DateReleased:     release.DateReleased,
				ReleasedQuantity: release.ReleasedQuantity,
				CanceledQuantity: release.CanceledQuantity,
				ShippedQuantity:  release.ShippedQuantity,
			})
		}

		oapiOrder.Items = append(oapiOrder.Items, oapiItem)
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
			HasStock: product.HasStock,
			IsActive: product.IsActive,
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

func mapQuoteStatus(status domain.QuoteStatus) oas.QuoteStatus {
	switch status {
	case domain.QuoteStatusProcessed:
		return oas.QuoteStatusProcessed
	case domain.QuoteStatusCancelled:
		return oas.QuoteStatusCancelled
	case domain.QuoteStatusPending:
		return oas.QuoteStatusPending
	case domain.QuoteStatusClosed:
		return oas.QuoteStatusClosed
	case domain.QuoteStatusUnspecified:
		return oas.QuoteStatusUnspecified
	default:
		return oas.QuoteStatusUnspecified
	}
}

func mapOrderItemDisposition(disposition domain.OrderItemDisposition) oas.OrderItemDisposition {
	switch disposition {
	case domain.OrderItemDispositionBackOrder:
		return oas.OrderItemDispositionBackorder
	case domain.OrderItemDispositionCancel:
		return oas.OrderItemDispositionCancel
	case domain.OrderItemDispositionDirectShip:
		return oas.OrderItemDispositionDirectShip
	case domain.OrderItemDispositionFuture:
		return oas.OrderItemDispositionFuture
	case domain.OrderItemDispositionHold:
		return oas.OrderItemDispositionHold
	case domain.OrderItemDispositionMultistageProcess:
		return oas.OrderItemDispositionMultistageProcess
	case domain.OrderItemDispositionProductionOrder:
		return oas.OrderItemDispositionProductionOrder
	case domain.OrderItemDispositionSpecialOrder:
		return oas.OrderItemDispositionSpecialOrder
	case domain.OrderItemDispositionTransfer:
		return oas.OrderItemDispositionTransfer
	default:
		return oas.OrderItemDispositionUnspecified
	}
}

func mapInvoiceAdjustmentType(adjustmentType domain.InvoiceAdjustmentType) oas.InvoiceAdjustmentType {
	switch adjustmentType {
	case domain.InvoiceAdjustmentTypeDebitMemo:
		return oas.InvoiceAdjustmentTypeDebitMemo
	case domain.InvoiceAdjustmentTypeCreditMemo:
		return oas.InvoiceAdjustmentTypeCreditMemo
	case domain.InvoiceAdjustmentTypeBadDebtWriteOff:
		return oas.InvoiceAdjustmentTypeBadDebtWriteOff
	case domain.InvoiceAdjustmentTypeBadDebtRecovery:
		return oas.InvoiceAdjustmentTypeBadDebtRecovery
	case domain.InvoiceAdjustmentTypeInvoice:
		return oas.InvoiceAdjustmentTypeInvoice
	default:
		return oas.InvoiceAdjustmentTypeUnspecified
	}
}
