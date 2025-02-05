package config

type Services struct {
	CatalogUrl  string `koanf:"catalog_url"`
	CustomerUrl string `koanf:"customer_url"`
	OrderUrl    string `koanf:"order_url"`
	SearchUrl   string `koanf:"search_url"`
}
