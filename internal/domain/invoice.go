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
