// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GetOrder implements getOrder operation.
//
// Get an order by ID.
//
// GET /orders/{id}
func (UnimplementedHandler) GetOrder(ctx context.Context, params GetOrderParams) (r *GetOrderOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetProduct implements getProduct operation.
//
// Get a product by ID.
//
// GET /products/{id}
func (UnimplementedHandler) GetProduct(ctx context.Context, params GetProductParams) (r GetProductRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListBranchOrders implements listBranchOrders operation.
//
// Get all orders for a customer branch.
//
// GET /orders
func (UnimplementedHandler) ListBranchOrders(ctx context.Context, params ListBranchOrdersParams) (r []Order, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCustomerBranches implements listCustomerBranches operation.
//
// Get available branches for customer.
//
// GET /account/branches
func (UnimplementedHandler) ListCustomerBranches(ctx context.Context, params ListCustomerBranchesParams) (r []Branch, _ error) {
	return r, ht.ErrNotImplemented
}

// ListOrderInvoices implements listOrderInvoices operation.
//
// Get invoices for an order.
//
// GET /orders/{id}/invoices
func (UnimplementedHandler) ListOrderInvoices(ctx context.Context, params ListOrderInvoicesParams) (r []InvoiceSimplified, _ error) {
	return r, ht.ErrNotImplemented
}

// ListOrderShipments implements listOrderShipments operation.
//
// Get shipments for an order.
//
// GET /orders/{id}/shipments
func (UnimplementedHandler) ListOrderShipments(ctx context.Context, params ListOrderShipmentsParams) (r []ShipmentSimplified, _ error) {
	return r, ht.ErrNotImplemented
}

// SearchProducts implements searchProducts operation.
//
// Search for products.
//
// POST /search/products
func (UnimplementedHandler) SearchProducts(ctx context.Context, req *SearchProductsReq) (r *SearchProductResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SetActiveBranch implements setActiveBranch operation.
//
// Set active branch for current user.
//
// PUT /account/branch
func (UnimplementedHandler) SetActiveBranch(ctx context.Context, req *SetActiveBranchReq) (r SetActiveBranchRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
