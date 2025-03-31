package domain

import "time"

type InvoiceSummary struct {
	Id           string
	OrderId      string
	Total        float64
	DateInvoiced time.Time
}
