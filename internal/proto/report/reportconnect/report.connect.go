// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: report.proto

package reportconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	report "github.com/materials-resources/store-api/internal/proto/report"
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
	// ReportServiceName is the fully-qualified name of the ReportService service.
	ReportServiceName = "report.v1.ReportService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ReportServiceGetInvoiceProcedure is the fully-qualified name of the ReportService's GetInvoice
	// RPC.
	ReportServiceGetInvoiceProcedure = "/report.v1.ReportService/GetInvoice"
	// ReportServiceGetPackingListProcedure is the fully-qualified name of the ReportService's
	// GetPackingList RPC.
	ReportServiceGetPackingListProcedure = "/report.v1.ReportService/GetPackingList"
)

// ReportServiceClient is a client for the report.v1.ReportService service.
type ReportServiceClient interface {
	GetInvoice(context.Context, *connect.Request[report.GetInvoiceRequest]) (*connect.ServerStreamForClient[report.GetInvoiceResponse], error)
	GetPackingList(context.Context, *connect.Request[report.GetPackingListRequest]) (*connect.ServerStreamForClient[report.GetPackingListResponse], error)
}

// NewReportServiceClient constructs a client for the report.v1.ReportService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReportServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ReportServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	reportServiceMethods := report.File_report_proto.Services().ByName("ReportService").Methods()
	return &reportServiceClient{
		getInvoice: connect.NewClient[report.GetInvoiceRequest, report.GetInvoiceResponse](
			httpClient,
			baseURL+ReportServiceGetInvoiceProcedure,
			connect.WithSchema(reportServiceMethods.ByName("GetInvoice")),
			connect.WithClientOptions(opts...),
		),
		getPackingList: connect.NewClient[report.GetPackingListRequest, report.GetPackingListResponse](
			httpClient,
			baseURL+ReportServiceGetPackingListProcedure,
			connect.WithSchema(reportServiceMethods.ByName("GetPackingList")),
			connect.WithClientOptions(opts...),
		),
	}
}

// reportServiceClient implements ReportServiceClient.
type reportServiceClient struct {
	getInvoice     *connect.Client[report.GetInvoiceRequest, report.GetInvoiceResponse]
	getPackingList *connect.Client[report.GetPackingListRequest, report.GetPackingListResponse]
}

// GetInvoice calls report.v1.ReportService.GetInvoice.
func (c *reportServiceClient) GetInvoice(ctx context.Context, req *connect.Request[report.GetInvoiceRequest]) (*connect.ServerStreamForClient[report.GetInvoiceResponse], error) {
	return c.getInvoice.CallServerStream(ctx, req)
}

// GetPackingList calls report.v1.ReportService.GetPackingList.
func (c *reportServiceClient) GetPackingList(ctx context.Context, req *connect.Request[report.GetPackingListRequest]) (*connect.ServerStreamForClient[report.GetPackingListResponse], error) {
	return c.getPackingList.CallServerStream(ctx, req)
}

// ReportServiceHandler is an implementation of the report.v1.ReportService service.
type ReportServiceHandler interface {
	GetInvoice(context.Context, *connect.Request[report.GetInvoiceRequest], *connect.ServerStream[report.GetInvoiceResponse]) error
	GetPackingList(context.Context, *connect.Request[report.GetPackingListRequest], *connect.ServerStream[report.GetPackingListResponse]) error
}

// NewReportServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReportServiceHandler(svc ReportServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	reportServiceMethods := report.File_report_proto.Services().ByName("ReportService").Methods()
	reportServiceGetInvoiceHandler := connect.NewServerStreamHandler(
		ReportServiceGetInvoiceProcedure,
		svc.GetInvoice,
		connect.WithSchema(reportServiceMethods.ByName("GetInvoice")),
		connect.WithHandlerOptions(opts...),
	)
	reportServiceGetPackingListHandler := connect.NewServerStreamHandler(
		ReportServiceGetPackingListProcedure,
		svc.GetPackingList,
		connect.WithSchema(reportServiceMethods.ByName("GetPackingList")),
		connect.WithHandlerOptions(opts...),
	)
	return "/report.v1.ReportService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ReportServiceGetInvoiceProcedure:
			reportServiceGetInvoiceHandler.ServeHTTP(w, r)
		case ReportServiceGetPackingListProcedure:
			reportServiceGetPackingListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedReportServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReportServiceHandler struct{}

func (UnimplementedReportServiceHandler) GetInvoice(context.Context, *connect.Request[report.GetInvoiceRequest], *connect.ServerStream[report.GetInvoiceResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("report.v1.ReportService.GetInvoice is not implemented"))
}

func (UnimplementedReportServiceHandler) GetPackingList(context.Context, *connect.Request[report.GetPackingListRequest], *connect.ServerStream[report.GetPackingListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("report.v1.ReportService.GetPackingList is not implemented"))
}
