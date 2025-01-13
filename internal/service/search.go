package service

import (
	"connectrpc.com/connect"
	"context"
	searchv1 "github.com/materials-resources/customer-api/internal/grpc-client/search"
	"github.com/materials-resources/customer-api/internal/grpc-client/search/searchconnect"
	"github.com/materials-resources/customer-api/internal/oas"
	"net/http"
)

type Search struct {
	Client searchconnect.SearchServiceClient
}

func NewSearchService() *Search {
	return &Search{
		Client: searchconnect.NewSearchServiceClient(http.DefaultClient,
			"http://localhost:8081", connect.WithGRPC()),
	}
}

func (s *Search) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (*oas.SearchProductsOK, error) {
	pbReq := &searchv1.SearchProductsRequest{
		Query:   req.GetQuery().Or(""),
		Page:    int32(req.GetPage().Or(1)),
		Filters: make(map[string]*searchv1.Filter),
	}

	for name, values := range req.GetFilters().Or(make(oas.SearchProductsReqFilters)) {
		pbReq.Filters[name] = &searchv1.Filter{
			Values: values,
		}
	}

	pbRes, err := s.Client.SearchProducts(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	response := oas.SearchProductsOK{
		Products: make([]oas.Product, 0), Metadata: oas.PageMetadata{
			TotalPages:   int(pbRes.Msg.GetPageMetadata().GetTotalPages()),
			TotalRecords: int(pbRes.Msg.GetPageMetadata().GetTotalRecords()),
		},
		Aggregations: make(map[string][]oas.Bucket),
	}

	for fbName, pbFb := range pbRes.Msg.GetAggregations().GetFieldBuckets() {
		var buckets []oas.Bucket
		for _, pbBucket := range pbFb.Aggregations {
			buckets = append(buckets, oas.Bucket{
				Value: pbBucket.GetName(),
				Count: int(pbBucket.GetCount()),
			})
		}
		response.Aggregations[fbName] = buckets
	}

	for _, pbProduct := range pbRes.Msg.GetResults() {
		response.Products = append(response.Products, oas.Product{
			ID:          pbProduct.GetBase().GetUid(),
			Sn:          pbProduct.GetBase().GetSerialNumber(),
			Name:        pbProduct.GetBase().GetName(),
			Description: oas.NewOptString(pbProduct.GetBase().GetDescription()),
			ImageURL:    oas.OptString{},
		})
	}

	return &response, nil

}
