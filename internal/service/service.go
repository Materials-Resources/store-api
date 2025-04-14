package service

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"github.com/materials-resources/store-api/app"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	Order    *Order
	Catalog  *CatalogService
	Search   *Search
	Customer *CustomerService
	Report   *ReportService
	Billing  *BillingService
}

func NewService(a *app.App) Service {

	return Service{
		Order:    NewOrderService(a),
		Catalog:  NewCatalogService(a),
		Search:   NewSearchService(a),
		Customer: NewCustomerService(a),
		Report:   NewReportService(a),
		Billing:  NewBillingService(a),
	}
}

func newInterceptor(tp trace.TracerProvider, mp metric.MeterProvider, p propagation.TextMapPropagator) (connect.Interceptor, error) {
	return otelconnect.NewInterceptor(
		otelconnect.WithTracerProvider(tp),
		otelconnect.WithMeterProvider(mp),
		otelconnect.WithPropagator(p),
	)
}
