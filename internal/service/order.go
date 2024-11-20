package service

import (
	"customer-api/client/microservices/proto/order/v1"
	"google.golang.org/grpc"
)

type Order struct {
	Client order.OrderServiceClient
}

func NewOrderService() *Order {
	conn, err := grpc.Dial("localhost:50058", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &Order{
		Client: order.NewOrderServiceClient(conn),
	}
}
