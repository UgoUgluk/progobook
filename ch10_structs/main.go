package main

import "fmt"

func main() {
	fmt.Println("Hello, Structs")

	type Product struct {
		name, category string
		price          float64
		int            //built-in field
	}

	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
	//only positional arguments
	var kayak2 = Product{"Kayak2", "Watersports2", 375.00, 1}
	fmt.Println("Name:", kayak2.name)
	fmt.Println("Category:", kayak2.category)
	fmt.Println("Price:", kayak2.price)
	fmt.Println(kayak)
}
