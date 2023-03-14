package store

//SpecialDeal expands product for sale
type SpecialDeal struct {
	Name string
	*Product
	price float64
}

//NewSpecialDeal create special deal
func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
	return &SpecialDeal{name, p, p.price - discount}
}

//GetDetails get data of special deal
func (deal *SpecialDeal) GetDetails() (string, float64, float64) {
	return deal.Name, deal.price, deal.Price(0)
}

//Price get price for special deal
func (deal *SpecialDeal) Price(taxRate float64) float64 {
	return deal.price
}
