package handler

import (
	"context"
	"fmt"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/service"
	"github.com/materials-resources/customer-api/internal/zitadel"
)

func NewHandler(service service.Service) Handler {
	z, err := zitadel.NewZitadelClient()
	if err != nil {
		panic(err)
	}
	return Handler{
		z:       z,
		service: service,
	}
}

type Handler struct {
	service service.Service
	z       *zitadel.Client
}

func (h Handler) ListOrders(ctx context.Context, params oas.ListOrdersParams) ([]oas.OrderSummary, error) {
	//TODO implement me
	panic("implement me")
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
	fmt.Println(req.GetDateRequested())
	fmt.Println(req.GetItems())

	response := oas.CreateQuoteCreated{
		QuoteID: oas.OptString{Value: "1234567890", Set: true},
		Status:  oas.OptString{Value: "PENDING", Set: true},
	}

	return &response, nil
}

func (h Handler) GetOrder(ctx context.Context, params oas.GetOrderParams) (oas.GetOrderRes, error) {

	res, err := h.service.Order.GetOrder(ctx, params)

	return res, err
}

func (h Handler) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	res, err := h.service.Catalog.GetProduct(ctx, params)
	return res, err
}

func (h Handler) ListCustomerBranches(ctx context.Context, params oas.ListCustomerBranchesParams) (oas.ListCustomerBranchesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListOrderInvoices(ctx context.Context, params oas.ListOrderInvoicesParams) (oas.ListOrderInvoicesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListOrderShipments(ctx context.Context, params oas.ListOrderShipmentsParams) (oas.ListOrderShipmentsRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (oas.SearchProductsRes, error) {
	return h.service.Search.SearchProducts(ctx, req)
}

var _ oas.Handler = (*Handler)(nil)
