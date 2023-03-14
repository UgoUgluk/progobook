package main

import "fmt"

// Expense represents a expense
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

// Account represents a account
type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {
	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println(
			"Expense:", expense.getName(),
			"Cost:", expense.getCost(true),
		)
	}

	fmt.Println("Total:", calcTotal(account.expenses))

}
