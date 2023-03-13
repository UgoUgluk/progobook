package main

import "fmt"

func printPriceSimple() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}

func printPrice(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func main() {
	fmt.Println("Hello, Functions")

	//simple function
	printPriceSimple()

	//function with parameters
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)
	//Go does not have optional or default parameters!
}
