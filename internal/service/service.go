package service

import "github.com/materials-resources/store-api/app"

type Service struct {
	Order    *Order
	Catalog  *CatalogService
	Search   *Search
	Customer *CustomerService
	Report   *ReportService
	Billing  *BillingService
}

func NewService(a *app.App) Service {

	return Service{
		Order:    NewOrderService(a),
		Catalog:  NewCatalogService(),
		Search:   NewSearchService(),
		Customer: NewCustomerService(),
		Report:   NewReportService(),
		Billing:  NewBillingService(),
	}
}
