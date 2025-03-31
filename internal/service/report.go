package service

import (
	"connectrpc.com/connect"
	"context"
	reportv1 "github.com/materials-resources/store-api/internal/proto/report"
	"github.com/materials-resources/store-api/internal/proto/report/reportconnect"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

type ReportService struct {
	client reportconnect.ReportServiceClient
}

func NewReportService() *ReportService {
	return &ReportService{
		client: reportconnect.NewReportServiceClient(http.DefaultClient,
			"http://localhost:8083", connect.WithGRPC()),
	}
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
