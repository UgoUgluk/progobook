package main

import (
	"fmt"
	_ "packages/data"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello, Packages and Modules")

	product := store.NewProduct("Kayak", "Watersports", 279)
	product2 := store.NewProduct("Kayak2", "Watersports2", 379)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product, *product2},
	}
	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))

	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + currencyFmt.ToCurrency(cart.GetTotal()))

}
