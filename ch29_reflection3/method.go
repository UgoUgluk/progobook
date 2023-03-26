package main

func (p Purchase) calcTax(taxRate float64) float64 {
	return p.Price * taxRate
}

// GetTotal get total price of purchase
func (p Purchase) GetTotal() float64 {
	return p.Price + p.calcTax(.20)
}
