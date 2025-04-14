// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// ContactUs implements contactUs operation.
	//
	// Send details regarding a contact inquiry.
	//
	// POST /contact
	ContactUs(ctx context.Context, req *ContactUsReq) (ContactUsRes, error)
	// CreateQuote implements createQuote operation.
	//
	// Create a new quote. The `customer_id` and `contact_id` are extracted from the provided
	// authentication token. Make sure to include a valid bearer token in the `Authorization` header.
	//
	// POST /account/quotes
	CreateQuote(ctx context.Context, req *CreateQuoteReq) (CreateQuoteRes, error)
	// GetActiveBranch implements getActiveBranch operation.
	//
	// Get active branch for user.
	//
	// GET /account/branches/active
	GetActiveBranch(ctx context.Context) (GetActiveBranchRes, error)
	// GetInvoiceReport implements getInvoiceReport operation.
	//
	// Get invoice report by ID.
	//
	// GET /account/invoices/{id}/report
	GetInvoiceReport(ctx context.Context, params GetInvoiceReportParams) (GetInvoiceReportRes, error)
	// GetOrder implements getOrder operation.
	//
	// Get an order by ID.
	//
	// GET /account/orders/{id}
	GetOrder(ctx context.Context, params GetOrderParams) (GetOrderRes, error)
	// GetPackingListReport implements getPackingListReport operation.
	//
	// Get packing list report by ID.
	//
	// GET /account/packinglist/{id}/report
	GetPackingListReport(ctx context.Context, params GetPackingListReportParams) (GetPackingListReportRes, error)
	// GetProduct implements getProduct operation.
	//
	// Get a product by ID.
	//
	// GET /products/{id}
	GetProduct(ctx context.Context, params GetProductParams) (GetProductRes, error)
	// GetQuote implements getQuote operation.
	//
	// Get quote by ID.
	//
	// GET /account/quotes/{id}
	GetQuote(ctx context.Context, params GetQuoteParams) (GetQuoteRes, error)
	// GetRecentPurchases implements getRecentPurchases operation.
	//
	// Get recent purchases for customer.
	//
	// GET /account/recent-purchases
	GetRecentPurchases(ctx context.Context, params GetRecentPurchasesParams) (GetRecentPurchasesRes, error)
	// ListCustomerBranches implements listCustomerBranches operation.
	//
	// Get available branches for customer.
	//
	// GET /account/branch
	ListCustomerBranches(ctx context.Context) (ListCustomerBranchesRes, error)
	// ListInvoices implements listInvoices operation.
	//
	// Get a list of invoices.
	//
	// GET /account/invoices
	ListInvoices(ctx context.Context, params ListInvoicesParams) (ListInvoicesRes, error)
	// ListOrders implements listOrders operation.
	//
	// Get a list of orders.
	//
	// GET /account/orders
	ListOrders(ctx context.Context, params ListOrdersParams) (ListOrdersRes, error)
	// ListQuotes implements listQuotes operation.
	//
	// Get a list of quotes.
	//
	// GET /account/quotes
	ListQuotes(ctx context.Context, params ListQuotesParams) (ListQuotesRes, error)
	// SearchProducts implements searchProducts operation.
	//
	// Search for products.
	//
	// POST /search/products
	SearchProducts(ctx context.Context, req *SearchProductsReq) (SearchProductsRes, error)
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
