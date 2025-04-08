package service

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/store-api/internal/domain"
	customerv1 "github.com/materials-resources/store-api/internal/proto/customer"
	"github.com/materials-resources/store-api/internal/proto/customer/customerconnect"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type CustomerService struct {
	Client customerconnect.CustomerServiceClient
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		Client: customerconnect.NewCustomerServiceClient(http.DefaultClient, "http://localhost:8082", connect.WithGRPC()),
	}
}

func (s *CustomerService) ListBranches(ctx context.Context, contactId string) ([]*domain.BranchSummary, error) {
	pbReq := customerv1.GetBranchesForContactRequest_builder{ContactId: proto.String(contactId)}.Build()

	pbRes, err := s.Client.GetBranchesForContact(ctx, connect.NewRequest(pbReq))

	if err != nil {
		return nil, err
	}

	var branches []*domain.BranchSummary

	for _, branch := range pbRes.Msg.GetBranches() {
		branches = append(branches, &domain.BranchSummary{
			Id:   branch.GetId(),
			Name: branch.GetName(),
		})
	}

	return branches, nil
}
