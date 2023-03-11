package main

import "fmt"

func main() {
	first := 100
	second := &first
	third := &second
	fmt.Println(first)
	fmt.Println(*second)
	fmt.Println(**third)
}
