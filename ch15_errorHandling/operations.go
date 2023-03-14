package main

// TotalPrice calc sum prices of group
func (slice ProductSlice) TotalPrice(category string) (total float64) {
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
		}
	}
	return
}
