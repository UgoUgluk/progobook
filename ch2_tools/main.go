package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

//PrintHello out Hello message
func PrintHello() {
	fmt.Println("Hello, Go")
}

//PrintNumber out Number message
func PrintNumber(number int) {
	fmt.Println(number)
}
