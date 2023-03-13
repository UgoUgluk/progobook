package main

import "fmt"

func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}
	fmt.Println("Product:", kayak.name, "Price:",
		kayak.price)
	fmt.Println("Service:", insurance.description, "Price:",
		insurance.monthlyFee*float64(insurance.durationMonths))
}
