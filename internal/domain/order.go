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
	DateOrdered   time.Time
	DateRequested time.Time
}
type Order struct {
	Id            string
	ContactId     string
	BranchId      string
	PurchaseOrder string
	Status        OrderStatus
	DateOrdered   time.Time
	DateRequested time.Time

	Items           []*OrderItem
	ShippingAddress Address

	Taker string
}

type OrderItemDisposition int

const (
	OrderItemDispositionUnspecified OrderItemDisposition = iota
	OrderItemDispositionBackOrder
	OrderItemDispositionCancel
	OrderItemDispositionDirectShip
	OrderItemDispositionFuture
	OrderItemDispositionHold
	OrderItemDispositionMultistageProcess
	OrderItemDispositionProductionOrder
	OrderItemDispositionSpecialOrder
	OrderItemDispositionTransfer
)

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

	Disposition OrderItemDisposition
	Releases    []*OrderItemRelease
}

type OrderItemRelease struct {
	DateReleased     time.Time
	ReleasedQuantity float64
	ShippedQuantity  float64
	CanceledQuantity float64
}

type OrderFilters struct {
	OrderId       string
	PurchaseOrder string
}
