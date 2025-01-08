package handler

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	orderv1 "github.com/materials-resources/customer-api/internal/grpc-client/order"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/service"
	"github.com/materials-resources/customer-api/internal/zitadel"
)

func NewHandler(service service.Service) Handler {
	z, err := zitadel.NewZitadelClient()
	if err != nil {
		panic(err)
	}
	return Handler{
		z:       z,
		service: service,
	}
}

type Handler struct {
	service service.Service
	z       *zitadel.Client
}

func (h Handler) SetActiveBranch(ctx context.Context, req *oas.SetActiveBranchReq) (oas.SetActiveBranchRes, error) {
	fmt.Println("changed branch")
	// check if user can access this branch
	// get user id from token

	// update the branch
	h.z.ChangeUserBranchId(ctx, "295379791043934934", req.GetBranchID())
	// return success

	return &oas.SetActiveBranchOK{}, nil
}

func (h Handler) CreateQuote(ctx context.Context, req *oas.CreateQuoteReq) (*oas.CreateQuoteCreated, error) {
	fmt.Println(req.GetRequestDate())
	fmt.Println(req.GetItems())

	response := oas.CreateQuoteCreated{
		QuoteID: oas.OptString{Value: "1234567890", Set: true},
		Status:  oas.OptString{Value: "PENDING", Set: true},
	}

	return &response, nil
}

func (h Handler) GetOrder(ctx context.Context, params oas.GetOrderParams) (*oas.GetOrderOK, error) {

	res, err := h.service.Order.Client.GetOrder(ctx, connect.NewRequest(&orderv1.GetOrderRequest{
		Id: params.ID,
	}))

	if err != nil {
		return nil, err
	}
	return &oas.GetOrderOK{
		Details: oas.Order{
			ID:        res.Msg.GetOrder().GetId(),
			OrderDate: res.Msg.GetOrder().GetDateCreated().AsTime().String(),
			Customer: oas.Customer{
				ID:   res.Msg.GetOrder().GetCustomer().GetId(),
				Name: res.Msg.GetOrder().GetCustomer().GetName(),
			},
			ContactID:            res.Msg.GetOrder().GetOrderDetails().GetContact().GetId(),
			ContactName:          res.Msg.GetOrder().GetOrderDetails().GetContact().GetFullName(),
			Taker:                res.Msg.GetOrder().GetOrderDetails().GetTaker(),
			PurchaseOrder:        res.Msg.GetOrder().GetPurchaseOrder(),
			DeliveryInstructions: res.Msg.GetOrder().GetOrderDetails().GetDeliveryInstructions(),
			ShippingAddress:      oas.Address{},
			Total:                0,
		},
	}, nil
}

func (h Handler) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	res, err := h.service.Catalog.GetProduct(ctx, params)
	return res, err
}

func (h Handler) ListBranchOrders(ctx context.Context, params oas.ListBranchOrdersParams) ([]oas.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListCustomerBranches(ctx context.Context, params oas.ListCustomerBranchesParams) ([]oas.Branch, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListOrderInvoices(ctx context.Context, params oas.ListOrderInvoicesParams) ([]oas.InvoiceSimplified, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ListOrderShipments(ctx context.Context, params oas.ListOrderShipmentsParams) ([]oas.ShipmentSimplified, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (*oas.SearchProductResponse, error) {
	return h.service.Search.SearchProducts(ctx, req)
}

var _ oas.Handler = (*Handler)(nil)
