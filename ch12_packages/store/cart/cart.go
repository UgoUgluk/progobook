package cart

import "packages/store"

// Cart describes a consumer cart
type Cart struct {
	CustomerName string
	Products     []store.Product
}

// GetTotal count sum of all cart items
func (cart *Cart) GetTotal() (total float64) {
	for _, p := range cart.Products {
		total += p.Price()
	}
	return
}
