package main

import "fmt"

// Product represents a product
type Product struct {
	name, category string
	price          float64
	//otherNames []string
}

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}
func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func main() {

	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}

	array := [1]StockLevel{
		{
			Product:   Product{"Kayak1", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)
	slice := []StockLevel{
		{
			Product:   Product{"Kayak2", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)
	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak3", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(kvp)

	//copy and link
	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p2 := p1
	p3 := &p1
	p1.name = "Original Kayak"
	fmt.Println("P1:", p1.name)    // P1: Original Kayak
	fmt.Println("P2:", p2.name)    //P2: Kayak
	fmt.Println("P3:", (*p3).name) //!(*p3)!   P3: Original Kayak

	//struct pointer
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax(&kayak)
	fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price", kayak.price)

	//function struct builder
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
		fmt.Println(*p)
	}
}
