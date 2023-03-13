package main

import "fmt"

func printPriceSimple() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}

func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func printPriceEmptyAttr(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func printSuppliersVar(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func printSuppliersVarEmpty(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func swapValuesLocal(first, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}
func swapValuesGlobal(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

func swapValuesByReturn(first, second int) (int, int) {
	return second, first
}

func calcDifTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcDifTax(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
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

	//function with empty parametr
	printPriceEmptyAttr("Kayak", 275, 0.2)
	printPriceEmptyAttr("Lifejacket", 48.95, 0.2)
	printPriceEmptyAttr("Soccer Ball", 19.50, 0.15)

	//function with slice parameter
	printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	printSuppliers("Lifejacket", []string{"Sail Safe Co"})

	//function with variable parameter(only last position)
	printSuppliersVar("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliersVar("Lifejacket", "Sail Safe Co")

	//function with empty variable parameter
	printSuppliersVarEmpty("Soccer Ball")

	//function which get slice(variable and not)
	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliersVar("Kayak", names...)
	printSuppliers("Slice Kayak", names)

	//function pointers
	//without
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValuesLocal(val1, val2)
	fmt.Println("After calling function", val1, val2)
	//with
	fmt.Println("Before calling function", val1, val2)
	swapValuesGlobal(&val1, &val2)
	fmt.Println("After calling function", val1, val2)

	//function with return
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		priceWithTax := calcTax(price)
		fmt.Println("Product: ", product, "Price:", priceWithTax)
	}

	//function returning two results
	fmt.Println("Before calling function", val1, val2)
	val1, val2 = swapValuesByReturn(val1, val2)
	fmt.Println("After calling function", val1, val2)

	//function returning current results + deffer
	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)
}
