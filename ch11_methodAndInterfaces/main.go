package main

import "fmt"

// Expense represents a expense
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

// Person represents a person
type Person struct {
	name, city string
}

func processItems(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println(
				"Service:", value.description,
				"Price:", value.monthlyFee*float64(value.durationMonths),
			)
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println(
				"Person Pointer:", value.name,
				"City:", value.city,
			)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	//get Service dinamic data by IF
	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println(
				"Service:", s.description,
				"Price:", s.monthlyFee*float64(s.durationMonths),
			)
		} else {
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}
	//get type dinamic data by switch
	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println(
				"Service:", value.description,
				"Price:", value.monthlyFee*float64(value.durationMonths),
			)
		case *Product:
			fmt.Println(
				"Product:", value.name,
				"Price:", value.price,
			)
		default:
			fmt.Println(
				"Expense:", expense.getName(),
				"Cost:", expense.getCost(true),
			)
		}
	}
	//empty interface
	fmt.Println("empty interface")
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println(
				"Service:", value.description,
				"Price:", value.monthlyFee*float64(value.durationMonths),
			)
		case Person:
			fmt.Println(
				"Person:", value.name,
				"City:", value.city,
			)
		case *Person:
			fmt.Println(
				"Person Pointer:", value.name,
				"City:", value.city,
			)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}

	//iterface - function attribute
	fmt.Println("function attribute")
	for _, item := range data {
		processItems(item)
	}

	//function with many iterface attributes
	fmt.Println("function with many iterface attributes")
	processItems(data...)

}
