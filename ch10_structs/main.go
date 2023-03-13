package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}

func main() {
	fmt.Println("Hello, Structs")

	type Product struct {
		name, category string
		price          float64
	}
	type Item struct {
		name     string
		category string
		price    float64
	}

	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}
	stockItem := StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.00},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count:", stockItem.count)
	fmt.Println("Alt Name:", stockItem.Alternate.name, stockItem.name)

	//change type and compare
	prod := Product{name: "Kayak", category: "Watersports",
		price: 275.00}
	item := Item{name: "Kayak", category: "Watersports",
		price: 275.00}
	fmt.Println("prod == item:", prod == Product(item))

	//anonymous structure types
	item2 := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod)
	writeName(item2)

	//anonymous structure: json
	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())

}
