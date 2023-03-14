package main

import "fmt"

//CalcStoreTotal  calc sum prices of all groups
func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		storeTotal += group.TotalPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

//TotalPrice calc sum prices of group
func (group ProductGroup) TotalPrice(category string) (total float64) {
	for _, p := range group {
		total += p.Price
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
