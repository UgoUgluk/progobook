package main

import (
	"fmt"
	//"strconv"
)

func main() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}
