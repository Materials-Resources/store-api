package service

type Service struct {
	Order   *Order
	Catalog *Catalog
	Search  *Search
}

func NewService() Service {

	return Service{
		Order:   NewOrderService(),
		Catalog: NewCatalogService(),
		Search:  NewSearchService(),
	}
}
