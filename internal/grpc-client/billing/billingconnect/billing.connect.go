// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: billing.proto

package billingconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	billing "github.com/materials-resources/store-api/internal/grpc-client/billing"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BillingServiceName is the fully-qualified name of the BillingService service.
	BillingServiceName = "billing.v1.BillingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BillingServiceGetInvoicesByOrderProcedure is the fully-qualified name of the BillingService's
	// GetInvoicesByOrder RPC.
	BillingServiceGetInvoicesByOrderProcedure = "/billing.v1.BillingService/GetInvoicesByOrder"
	// BillingServiceGetInvoiceProcedure is the fully-qualified name of the BillingService's GetInvoice
	// RPC.
	BillingServiceGetInvoiceProcedure = "/billing.v1.BillingService/GetInvoice"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	billingServiceServiceDescriptor                  = billing.File_billing_proto.Services().ByName("BillingService")
	billingServiceGetInvoicesByOrderMethodDescriptor = billingServiceServiceDescriptor.Methods().ByName("GetInvoicesByOrder")
	billingServiceGetInvoiceMethodDescriptor         = billingServiceServiceDescriptor.Methods().ByName("GetInvoice")
)

// BillingServiceClient is a client for the billing.v1.BillingService service.
type BillingServiceClient interface {
	// GetInvoicesByOrder returns all invoices for a given order
	GetInvoicesByOrder(context.Context, *connect.Request[billing.GetInvoicesByOrderRequest]) (*connect.Response[billing.GetInvoicesByOrderResponse], error)
	GetInvoice(context.Context, *connect.Request[billing.GetInvoiceRequest]) (*connect.Response[billing.GetInvoiceResponse], error)
}

// NewBillingServiceClient constructs a client for the billing.v1.BillingService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBillingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BillingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &billingServiceClient{
		getInvoicesByOrder: connect.NewClient[billing.GetInvoicesByOrderRequest, billing.GetInvoicesByOrderResponse](
			httpClient,
			baseURL+BillingServiceGetInvoicesByOrderProcedure,
			connect.WithSchema(billingServiceGetInvoicesByOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getInvoice: connect.NewClient[billing.GetInvoiceRequest, billing.GetInvoiceResponse](
			httpClient,
			baseURL+BillingServiceGetInvoiceProcedure,
			connect.WithSchema(billingServiceGetInvoiceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// billingServiceClient implements BillingServiceClient.
type billingServiceClient struct {
	getInvoicesByOrder *connect.Client[billing.GetInvoicesByOrderRequest, billing.GetInvoicesByOrderResponse]
	getInvoice         *connect.Client[billing.GetInvoiceRequest, billing.GetInvoiceResponse]
}

// GetInvoicesByOrder calls billing.v1.BillingService.GetInvoicesByOrder.
func (c *billingServiceClient) GetInvoicesByOrder(ctx context.Context, req *connect.Request[billing.GetInvoicesByOrderRequest]) (*connect.Response[billing.GetInvoicesByOrderResponse], error) {
	return c.getInvoicesByOrder.CallUnary(ctx, req)
}

// GetInvoice calls billing.v1.BillingService.GetInvoice.
func (c *billingServiceClient) GetInvoice(ctx context.Context, req *connect.Request[billing.GetInvoiceRequest]) (*connect.Response[billing.GetInvoiceResponse], error) {
	return c.getInvoice.CallUnary(ctx, req)
}

// BillingServiceHandler is an implementation of the billing.v1.BillingService service.
type BillingServiceHandler interface {
	// GetInvoicesByOrder returns all invoices for a given order
	GetInvoicesByOrder(context.Context, *connect.Request[billing.GetInvoicesByOrderRequest]) (*connect.Response[billing.GetInvoicesByOrderResponse], error)
	GetInvoice(context.Context, *connect.Request[billing.GetInvoiceRequest]) (*connect.Response[billing.GetInvoiceResponse], error)
}

// NewBillingServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBillingServiceHandler(svc BillingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	billingServiceGetInvoicesByOrderHandler := connect.NewUnaryHandler(
		BillingServiceGetInvoicesByOrderProcedure,
		svc.GetInvoicesByOrder,
		connect.WithSchema(billingServiceGetInvoicesByOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetInvoiceHandler := connect.NewUnaryHandler(
		BillingServiceGetInvoiceProcedure,
		svc.GetInvoice,
		connect.WithSchema(billingServiceGetInvoiceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/billing.v1.BillingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BillingServiceGetInvoicesByOrderProcedure:
			billingServiceGetInvoicesByOrderHandler.ServeHTTP(w, r)
		case BillingServiceGetInvoiceProcedure:
			billingServiceGetInvoiceHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBillingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBillingServiceHandler struct{}

func (UnimplementedBillingServiceHandler) GetInvoicesByOrder(context.Context, *connect.Request[billing.GetInvoicesByOrderRequest]) (*connect.Response[billing.GetInvoicesByOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("billing.v1.BillingService.GetInvoicesByOrder is not implemented"))
}

func (UnimplementedBillingServiceHandler) GetInvoice(context.Context, *connect.Request[billing.GetInvoiceRequest]) (*connect.Response[billing.GetInvoiceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("billing.v1.BillingService.GetInvoice is not implemented"))
}
