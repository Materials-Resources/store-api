package domain

type Product struct {
	Id                     string
	Sn                     string
	Name                   string
	Description            string
	ProductGroupId         string
	ProductGroupName       string
	SalesUnitOfMeasurement UnitOfMeasurement
	IsActive               bool
	HasStock               bool
}
