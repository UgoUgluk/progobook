package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf(
			"getProductName: %v",
			Products[index].Name,
		)
	} else {
		err = fmt.Errorf("Error for index %v", index)
	}
	return
}

//Printfln printf line
func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scanln(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln(
			"Name: %v, Category: %v, Price: %.2f",
			name, category, price,
		)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
