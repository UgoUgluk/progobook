package main

import (
	"composition/store"
	"fmt"
)

func main() {
	fmt.Println("Hello, Composition")
	//products
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}
	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println(
			"Name:", p.Name,
			"Category:", p.Category,
			"Price:", p.Price(0.2),
		)
	}
	//boats
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		fmt.Println(
			"Conventional:", b.Product.Name,
			"Direct:", b.Name,
		)
	}

	// rentalboats
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println(
			"Rental Boat:", r.Name,
			"Rental Price:", r.Price(0.2),
			"Captain:", r.Captain,
		)
	}

	//special deal
	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)

	//OfferBundle
	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}
	bundle := OfferBundle{
		store.NewSpecialDeal("Weekend Special", kayak, 50),
		store.NewProduct("Lifrejacket", "Watersports", 48.95),
	}
	//fmt.Println("Price:", bundle.Price(0)) - error: ambiguous selector bundle.Price
	fmt.Println("Price:", bundle.SpecialDeal.Price(0))
	fmt.Println("Price:", bundle.Product.Price(0))
}
