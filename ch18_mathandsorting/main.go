package main

//"sort"

func main() {
	products := []Product{
		{"Kayak", 279},
		{"Soccer Ball", 19.50},
		{"Lifejacket", 49.95},
	}

	//sort by name
	Printfln("sort by name")
	ProductSlicesByName(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}

	//sort by price
	Printfln("sort by price")
	ProductSlices(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}

	//flex sort by name
	Printfln("flex sort by name")
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}

	//flex sort by price
	Printfln("flex sort by price")
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Price < p2.Price
	})
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
