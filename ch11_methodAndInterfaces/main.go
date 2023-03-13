package main

import "fmt"

// Product represents a product
type Product struct {
	name, category string
	price          float64
}

// ProductList represents a list of products
type ProductList []Product

// Supplier represents a supplier
type Supplier struct {
	name, city string
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}
func (product *Product) metodPrintDetails() {
	fmt.Println(
		"Name:", product.name,
		"Category:", product.category,
		"Price", product.calcTax(0.2, 100),
	)
}

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func (supplier *Supplier) printDetails() {
	fmt.Println(
		"Supplier:", supplier.name,
		"City:", supplier.city,
	)
}

func funcPrintDetails(product *Product) {
	fmt.Println(
		"Name:", product.name,
		"Category:", product.category,
		"Price", product.price,
	)
}

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		funcPrintDetails(p)
		p.metodPrintDetails()
	}
	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}

	kayak := &Product{"Kayak", "Watersports", 275}
	kayak.metodPrintDetails()

	//alies
	products2 := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for category, total := range products2.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

}
