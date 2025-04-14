package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/app"
	reportv1 "github.com/materials-resources/store-api/internal/proto/report"
	"github.com/materials-resources/store-api/internal/proto/report/reportconnect"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

type ReportService struct {
	client reportconnect.ReportServiceClient
}

func NewReportService(a *app.App) *ReportService {
	otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
	if err != nil {
		a.Logger.Fatal().Str("service", "report").Err(err).Msg("could not create otel interceptor")
	}
	return &ReportService{
		client: reportconnect.NewReportServiceClient(http.DefaultClient,
			a.Config.Services.ReportUrl, connect.WithInterceptors(otelInterceptor), connect.WithGRPC()),
	}
}

func (s *ReportService) GetPackingList(ctx context.Context, id string) (io.ReadCloser, error) {
	pbReq := reportv1.GetPackingListRequest_builder{Id: proto.String(id)}.Build()

	stream, err := s.client.GetPackingList(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()

	// Process the stream in a goroutine
	go func() {
		defer pw.Close()

		// Read chunks from the stream and write to the pipe
		for stream.Receive() {
			// Write the chunk to the pipe
			_, err = pw.Write(stream.Msg().GetContent())
			if err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()

	return pr, nil
}

func (s *ReportService) GetInvoice(ctx context.Context, id string) (io.ReadCloser, error) {
	pbReq := reportv1.GetInvoiceRequest_builder{Id: proto.String(id)}.Build()

	stream, err := s.client.GetInvoice(ctx, connect.NewRequest(pbReq))
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()

	// Process the stream in a goroutine
	go func() {
		defer pw.Close()

		// Read chunks from the stream and write to the pipe
		for stream.Receive() {
			// Write the chunk to the pipe
			_, err = pw.Write(stream.Msg().GetContent())
			if err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()

	return pr, nil
}
