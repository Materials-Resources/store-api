package app

import (
	"context"
	"github.com/materials-resources/store-api/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type Otel struct {
	tp  *trace.TracerProvider
	mp  *metric.MeterProvider
	tmp propagation.TextMapPropagator
}

func NewOtel(cfg config.Config) (*Otel, error) {
	ctx := context.Background()

	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	tp := newTracerProvider(exp, cfg.Telemetry.ServiceName, cfg.Env)
	mp := newMeterProvider()
	tmp := newTextMapPropagator()
	return &Otel{
		tmp: tmp,
		tp:  tp,
		mp:  mp,
	}, nil
}

func newTracerProvider(exp *otlptrace.Exporter, serviceName, environment string) *trace.TracerProvider {

	return trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				attribute.String("environment", environment),
			),
		),
	)
}

func newMeterProvider() *metric.MeterProvider {
	return metric.NewMeterProvider()
}

func newTextMapPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{})
}

func (o *Otel) GetTracerProvider() *trace.TracerProvider {
	return o.tp
}

func (o *Otel) GetMeterProvider() *metric.MeterProvider {
	return o.mp
}

func (o *Otel) GetTextMapPropagator() propagation.TextMapPropagator {

	return o.tmp
}
