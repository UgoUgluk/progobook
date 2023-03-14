package main

import "fmt"

// Expense represents a expense
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = &product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println(
		"Expense method result:",
		expense.getCost(false),
	)
	//compare
	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}
	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2) //false because different place of memory
	fmt.Println("e3 == e4", e3 == e4)
	fmt.Println(e1)
	fmt.Println(e2)
}
