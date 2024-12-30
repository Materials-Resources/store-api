package service

import (
	"connectrpc.com/connect"
	"github.com/materials-resources/customer-api/internal/grpc-client/order/orderconnect"
	"net/http"
)

type Order struct {
	Client orderconnect.OrderServiceClient
}

func NewOrderService() *Order {
	return &Order{
		Client: orderconnect.NewOrderServiceClient(http.DefaultClient,
			"http://localhost:50058",
			connect.WithGRPC()),
	}
}
