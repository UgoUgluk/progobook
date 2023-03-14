package store

// Product describes an item for sale
type Product struct {
	Name, Category string
	price          float64
}

// NewProduct creates a new product
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Price calc price
func (p *Product) Price(taxRate float64) float64 {
	return p.price + (p.price * taxRate)
}

//Describable interface for product
type Describable interface {
	GetName() string
	GetCategory() string
	ItemForSale
}

//GetName return name
func (p *Product) GetName() string {
	return p.Name
}

//GetCategory return category
func (p *Product) GetCategory() string {
	return p.Category
}
