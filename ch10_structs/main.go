package main

import "fmt"

// Product represents a product
type Product struct {
	name, category string
	price          float64
	*Supplier
}

// Supplier represents a supplier
type Supplier struct {
	name, city string
}

func newProduct(
	name, category string,
	price float64,
	supplier *Supplier,
) *Product {
	return &Product{name, category, price - 10, supplier}
}
func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

func main() {
	acme := &Supplier{"Acme Co", "New York"}
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	//copy
	p1 := newProduct("Kayak", "Watersports", 275, acme)
	p2 := *p1
	p3 := copyProduct(p1)
	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"
	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}
	//manual copy
	for _, p := range []Product{*p1, p3} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	//null error

	//var prod Product
	var prod Product = Product{Supplier: &Supplier{}}

	var prodPtr *Product
	fmt.Println("Value:", prod.name, prod.category,
		prod.price,
		prod.Supplier.name) //if var prod Product  - Supplier == nil, prod.Supplier.name - error
	fmt.Println("Pointer:", prodPtr)

}
