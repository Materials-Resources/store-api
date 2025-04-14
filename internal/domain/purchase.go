package domain

type PurchaseSummary struct {
	ProductId          string
	ProductSn          string
	ProductName        string
	ProductDescription string
	UnitType           string
	OrderedQuantity    float64
}
