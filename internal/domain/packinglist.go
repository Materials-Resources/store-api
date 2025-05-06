package domain

import "time"

type PackingListSummary struct {
	InvoiceId    string
	OrderId      string
	DateInvoiced time.Time
}

type PackingList struct {
	InvoiceId    string
	OrderId      string
	CustomerId   string
	BranchId     string
	DateInvoiced time.Time
}
