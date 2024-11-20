package service

type Service struct {
	Order   *Order
	Catalog *Catalog
}

func NewService() Service {

	return Service{
		Order:   NewOrderService(),
		Catalog: NewCatalogService(),
	}
}
