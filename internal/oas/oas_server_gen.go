// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateQuote implements createQuote operation.
	//
	// Create a new quote. The `customer_id` and `contact_id` are extracted from the provided
	// authentication token. Make sure to include a valid bearer token in the `Authorization` header.
	//
	// POST /account/quotes
	CreateQuote(ctx context.Context, req *CreateQuoteReq) (*CreateQuoteCreated, error)
	// GetOrder implements getOrder operation.
	//
	// Get an order by ID.
	//
	// GET /orders/{id}
	GetOrder(ctx context.Context, params GetOrderParams) (*GetOrderOK, error)
	// GetProduct implements getProduct operation.
	//
	// Get a product by ID.
	//
	// GET /products/{id}
	GetProduct(ctx context.Context, params GetProductParams) (GetProductRes, error)
	// ListBranchOrders implements listBranchOrders operation.
	//
	// Get all orders for a customer branch.
	//
	// GET /orders
	ListBranchOrders(ctx context.Context, params ListBranchOrdersParams) ([]Order, error)
	// ListCustomerBranches implements listCustomerBranches operation.
	//
	// Get available branches for customer.
	//
	// GET /account/branches
	ListCustomerBranches(ctx context.Context, params ListCustomerBranchesParams) ([]Branch, error)
	// ListOrderInvoices implements listOrderInvoices operation.
	//
	// Get invoices for an order.
	//
	// GET /orders/{id}/invoices
	ListOrderInvoices(ctx context.Context, params ListOrderInvoicesParams) ([]InvoiceSimplified, error)
	// ListOrderShipments implements listOrderShipments operation.
	//
	// Get shipments for an order.
	//
	// GET /orders/{id}/shipments
	ListOrderShipments(ctx context.Context, params ListOrderShipmentsParams) ([]ShipmentSimplified, error)
	// SearchProducts implements searchProducts operation.
	//
	// Search for products.
	//
	// POST /search/products
	SearchProducts(ctx context.Context, req *SearchProductsReq) (*SearchProductResponse, error)
	// SetActiveBranch implements setActiveBranch operation.
	//
	// Set active branch for current user.
	//
	// PUT /account/branch
	SetActiveBranch(ctx context.Context, req *SetActiveBranchReq) (SetActiveBranchRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
