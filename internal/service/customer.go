package service

import (
	"connectrpc.com/connect"
	"github.com/materials-resources/store-api/internal/proto/customer/customerconnect"
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
