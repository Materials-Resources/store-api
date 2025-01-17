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
	CreateQuote(ctx context.Context, req *CreateQuoteReq) (CreateQuoteRes, error)
	// GetOrder implements getOrder operation.
	//
	// Get an order by ID.
	//
	// GET /account/orders/{id}
	GetOrder(ctx context.Context, params GetOrderParams) (GetOrderRes, error)
	// GetProduct implements getProduct operation.
	//
	// Get a product by ID.
	//
	// GET /products/{id}
	GetProduct(ctx context.Context, params GetProductParams) (GetProductRes, error)
	// ListCustomerBranches implements listCustomerBranches operation.
	//
	// Get available branches for customer.
	//
	// GET /account/branches
	ListCustomerBranches(ctx context.Context, params ListCustomerBranchesParams) (*ListCustomerBranchesOK, error)
	// ListOrders implements listOrders operation.
	//
	// Get a list of orders.
	//
	// GET /account/orders
	ListOrders(ctx context.Context, params ListOrdersParams) (*ListOrdersOK, error)
	// ListQuotes implements listQuotes operation.
	//
	// Get a list of quotes.
	//
	// GET /account/quotes
	ListQuotes(ctx context.Context, params ListQuotesParams) (*ListQuotesOK, error)
	// SearchProducts implements searchProducts operation.
	//
	// Search for products.
	//
	// POST /search/products
	SearchProducts(ctx context.Context, req *SearchProductsReq) (*SearchProductsOK, error)
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
