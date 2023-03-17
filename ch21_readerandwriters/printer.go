package main

import "fmt"

//Printfln printf to line
func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
