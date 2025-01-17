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

func (h Handler) ListQuotes(ctx context.Context, params oas.ListQuotesParams) (*oas.ListQuotesOK, error) {
	res, err := h.service.Order.ListQuotes(ctx, params)
	return res, err
}

func (h Handler) ListOrders(ctx context.Context, params oas.ListOrdersParams) (*oas.ListOrdersOK, error) {
	res, err := h.service.Order.ListOrders(ctx, params)
	return res, err

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

	res, err := h.service.Order.GetOrder(ctx, params)

	return res, err
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
