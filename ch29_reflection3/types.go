package main

//Product present a product
type Product struct {
	Name, Category string
	Price          float64
}

//Customer present a customer
type Customer struct {
	Name, City string
}

//Purchase present a purchase
type Purchase struct {
	Customer
	Product
	Total   float64
	taxRate float64
}
