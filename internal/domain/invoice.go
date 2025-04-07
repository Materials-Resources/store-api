package domain

import "time"

type InvoiceSummary struct {
	Id           string
	OrderId      string
	PaidAmount   float64
	TotalAmount  float64
	DateInvoiced time.Time
}
