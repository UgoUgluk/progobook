package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("printPrice  -- ", "Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func main() {
	fmt.Println("Hello, Function Types")

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {

		//function as variable
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:",
			totalPrice)

		//function as argument
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}

		//function as return value
		printPrice(product, price, selectCalculator(price))

	}

}
