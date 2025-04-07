package domain

import "time"

type PackingListSummary struct {
	InvoiceId    string
	OrderId      string
	DateInvoiced time.Time
}
