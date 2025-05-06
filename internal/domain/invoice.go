package domain

import "time"

type InvoiceAdjustmentType int

const (
	InvoiceAdjustmentTypeUnspecified InvoiceAdjustmentType = iota
	InvoiceAdjustmentTypeDebitMemo
	InvoiceAdjustmentTypeCreditMemo
	InvoiceAdjustmentTypeBadDebtWriteOff
	InvoiceAdjustmentTypeBadDebtRecovery
	InvoiceAdjustmentTypeInvoice
)

type InvoiceSummary struct {
	Id             string
	OrderId        string
	PaidAmount     float64
	TotalAmount    float64
	DateInvoiced   time.Time
	AdjustmentType InvoiceAdjustmentType
}

type Invoice struct {
	Id         string
	OrderId    string
	CustomerId string
	BranchId   string
	Totals     InvoiceTotals
}

type InvoiceTotals struct {
	SubTotal   float64
	AmountPaid float64
	AmountDue  float64
}
