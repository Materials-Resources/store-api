package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/internal/oas"
	searchv1 "github.com/materials-resources/store-api/internal/proto/search"
	"github.com/materials-resources/store-api/internal/proto/search/searchconnect"
	"google.golang.org/protobuf/proto"
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
	pbReq := searchv1.SearchProductsRequest_builder{
		Query:   proto.String(req.GetQuery().Or("")),
		Page:    proto.Int32(int32(req.GetPage().Or(1))),
		Filters: make(map[string]*searchv1.Filter),
	}

	for name, values := range req.GetFilters().Or(make(oas.SearchProductsReqFilters)) {
		pbReq.Filters[name] = searchv1.Filter_builder{
			Values: values,
		}.Build()
	}

	pbRes, err := s.Client.SearchProducts(ctx, connect.NewRequest(pbReq.Build()))

	if err != nil {
		return nil, err
	}

	response := oas.SearchProductsOK{
		Products: make([]oas.Product, 0), Metadata: oas.PageMetadata{
			TotalPages:   int(pbRes.Msg.GetPageMetadata().GetTotalPages()),
			TotalRecords: int(pbRes.Msg.GetPageMetadata().GetTotalRecords()),
		},
		Aggregations: make([]oas.Aggregation, 0),
	}

	for _, aggregationPb := range pbRes.Msg.GetAggregations() {
		aggregation := oas.Aggregation{}
		switch aggregationPb.WhichAggregationType() {
		case searchv1.Aggregation_TermsAggregation_case:

			termsAggregation := oas.TermsAggregation{
				FieldName: aggregationPb.GetTermsAggregation().GetFieldName(),
			}

			for _, term := range aggregationPb.GetTermsAggregation().GetBuckets() {
				termsAggregation.Buckets = append(termsAggregation.Buckets, oas.TermsAggregationBucket{
					Key:   term.GetKey(),
					Count: int(term.GetCount()),
				})
			}

			aggregation.SetTermsAggregation(termsAggregation)
		}
		response.Aggregations = append(response.Aggregations, aggregation)
	}

	for _, pbProduct := range pbRes.Msg.GetProducts() {
		response.Products = append(response.Products, oas.Product{
			ID:          pbProduct.GetId(),
			Sn:          pbProduct.GetSn(),
			Name:        pbProduct.GetName(),
			Description: pbProduct.GetDescription(),
			ImageURL:    oas.OptString{},
		})
	}

	return &response, nil

}
