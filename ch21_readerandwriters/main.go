package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}

	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)

	encoder.Encode(m)

	encoder.Encode(Kayak)

	encoder.Encode(&dp)

	dp2 := DiscountedProduct{Discount: 10.50}
	encoder.Encode(&dp2)

	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)

	fmt.Print(writer.String())

}
