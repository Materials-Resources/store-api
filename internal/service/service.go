package service

import "github.com/materials-resources/customer-api/app"

type Service struct {
	Order    *Order
	Catalog  *CatalogService
	Search   *Search
	Customer *CustomerService
}

func NewService(a *app.App) Service {

	return Service{
		Order:    NewOrderService(a),
		Catalog:  NewCatalogService(),
		Search:   NewSearchService(),
		Customer: NewCustomerService(),
	}
}
