package domain

import "time"

const (
	QuoteStatusProcessed QuoteStatus = iota
	QuoteStatusPending
	QuoteStatusCancelled
	QuoteStatusClosed
	QuoteStatusUnspecified
)

type QuoteStatus int

type QuoteFilters struct {
	Id            string
	PurchaseOrder string
}
type QuoteSummary struct {
	Id            string
	BranchId      string
	ContactId     string
	PurchaseOrder string
	Status        QuoteStatus
	DateCreated   time.Time
	DateExpires   time.Time
}
type Quote struct {
	Id            string
	BranchId      string
	ContactId     string
	PurchaseOrder string
	Status        QuoteStatus
	DateCreated   time.Time
	DateExpires   time.Time

	Items []*QuoteItem
}

type QuoteItem struct {
	ProductId         string
	ProductSn         string
	ProductName       string
	CustomerProductSn string
	OrderedQuantity   float64
	UnitPrice         float64
	UnitType          UnitOfMeasurement
	TotalPrice        float64
}
