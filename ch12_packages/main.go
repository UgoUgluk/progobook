package main

import (
	"fmt"
	"packages/store"
)

func main() {
	fmt.Println("Hello, Packages and Modules")

	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
}
