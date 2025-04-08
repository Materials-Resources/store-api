package domain

import "time"

const (
	OrderStatusCompleted OrderStatus = iota
	OrderStatusPendingApproval
	OrderStatusApproved
	OrderStatusCancelled
	OrderStatusUnspecified
)

type OrderStatus int
type OrderSummary struct {
	Id            string
	ContactId     string
	BranchId      string
	PurchaseOrder string
	Status        OrderStatus
	DateCreated   time.Time
	DateRequested time.Time
}
type Order struct {
	Id            string
	ContactId     string
	BranchId      string
	PurchaseOrder string
	Status        OrderStatus
	DateCreated   time.Time
	DateRequested time.Time

	Items           []*OrderItem
	ShippingAddress Address

	Taker string
}

type OrderItem struct {
	ProductId         string
	ProductSn         string
	ProductName       string
	CustomerProductSn string
	OrderedQuantity   float64
	ShippedQuantity   float64
	RemainingQuantity float64
	UnitType          UnitOfMeasurement
	UnitPrice         float64
	TotalPrice        float64
}
