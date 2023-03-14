package main

import "strconv"

//Product presente a product
type Product struct {
	Name, Category string
	Price          float64
}

//ProductSlice type of product slice
type ProductSlice []*Product

//Products create slice of product
var Products = ProductSlice{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

//ToCurrency presente a price in currency
func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}
