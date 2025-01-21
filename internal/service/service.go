package service

type Service struct {
	Order    *Order
	Catalog  *CatalogService
	Search   *Search
	Customer *CustomerService
}

func NewService() Service {

	return Service{
		Order:    NewOrderService(),
		Catalog:  NewCatalogService(),
		Search:   NewSearchService(),
		Customer: NewCustomerService(),
	}
}
