package handler

import (
	"context"
	"customer-api/client/microservices/proto/catalog/v1"
	"customer-api/client/microservices/proto/order/v1"
	"customer-api/internal/oas"
	"customer-api/internal/service"
	"customer-api/internal/zitadel"
	"fmt"
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

func (h Handler) GetOrder(ctx context.Context, params oas.GetOrderParams) (*oas.GetOrderOK, error) {

	res, err := h.service.Order.Client.GetOrder(ctx, &order.GetOrderRequest{
		Id: params.ID,
	})

	if err != nil {
		return nil, err
	}
	return &oas.GetOrderOK{
		Details: oas.Order{
			ID:        res.Order.Id,
			OrderDate: res.GetOrder().GetDateCreated().AsTime().String(),
			Customer: oas.Customer{
				ID:   res.GetOrder().GetCustomer().GetId(),
				Name: res.GetOrder().GetCustomer().GetName(),
			},
			ContactID:            res.GetOrder().GetOrderDetails().GetContact().GetId(),
			ContactName:          res.GetOrder().GetOrderDetails().GetContact().GetFullName(),
			Taker:                res.GetOrder().GetOrderDetails().GetTaker(),
			PurchaseOrder:        res.GetOrder().GetPurchaseOrder(),
			DeliveryInstructions: res.GetOrder().GetOrderDetails().GetDeliveryInstructions(),
			ShippingAddress:      oas.Address{},
			Total:                0,
		},
	}, nil
}

func (h Handler) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	res, err := h.service.Catalog.Client.GetProduct(ctx, &catalog.GetProductRequest{
		ProductUid: params.ID,
	})

	if err != nil {
		return nil, err
	}

	return &oas.GetProductOK{
		Details: oas.Product{
			ID:          res.GetProduct().GetUid(),
			Sn:          res.GetProduct().GetSn(),
			Name:        res.GetProduct().GetName(),
			Description: oas.OptString{Value: res.GetProduct().GetDescription(), Set: true},
			ImageURL:    oas.OptString{},
		},
	}, nil
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

func (h Handler) SearchProducts(ctx context.Context, req *oas.SearchProductsReq) (*oas.ProductSearchResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ oas.Handler = (*Handler)(nil)
