package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, Collections")

	//maps
	products := make(map[string]float64, 10)
	products["Kayak"] = 279
	products["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])

	//check 0 value
	products2 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}

	if value, ok := products2["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	//delete
	fmt.Println(products2)
	delete(products2, "Hat")
	fmt.Println(products2)

	//iterate
	products2["Hat"] = 0
	for key, value := range products2 {
		fmt.Println("Key:", key, "Value:", value)
	}

	//sort
	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}

	//string as array
	var price string = "€48.95"
	var currency byte = price[0]
	var currency2 string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	fmt.Println("Currency2:", currency2)
	fmt.Println("Length:", len(price))
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}

}
