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
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
	fmt.Printf("Product: %v, Price: $%4.2f\n", Kayak.Name, Kayak.Price)

	name, _ := getProductName(1)
	fmt.Println(name)
	_, err := getProductName(10)
	fmt.Println(err.Error())

	Printfln("Value: %v", Kayak)
	Printfln("Fields+Value: %+v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	//number
	number := 250
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)

	//float
	float := 279.00
	Printfln("Decimalless with exponent: %b", float)
	Printfln("Decimal with exponent: %e", float)
	Printfln("Decimal without exponent: %f", float)
	Printfln("Hexadecimal: %x, %X", float, float)
	Printfln("Decimal without exponent: >>%8.2f<<", float)
	Printfln("Sign: >>%+.2f<<", float)
	Printfln("Zeros for Padding: >>%010.2f<<", float)
	Printfln("Right Padding: >>%-8.2f<<", float)

	//string
	name2 := "Kayak"
	Printfln("String: %s", name2)
	Printfln("Character: %c", []rune(name2)[0])
	Printfln("Unicode: %U", []rune(name2)[0])

	//bool
	Printfln("Bool: %t", len(name2) > 1)
	Printfln("Bool: %t", len(name2) > 100)

	//Pointer
	Printfln("Pointer: %p", &name2)

}
