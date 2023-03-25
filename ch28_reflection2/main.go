package main

import (
	"reflect"
	"strings"
	// "fmt"
)

var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val interface{}) {
	elemValue := reflect.ValueOf(val)
	if elemValue.Type() == stringPtrType {
		upperStr := strings.ToUpper(elemValue.Elem().String())
		if elemValue.Elem().CanSet() {
			elemValue.Elem().SetString(upperStr)
		}
	}
}

func createPointerType(t reflect.Type) reflect.Type {
	return reflect.PtrTo(t)
}
func followPointerType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func main() {
	name := "Alice"
	t := reflect.TypeOf(name)
	Printfln("Original Type: %v", t)
	pt := createPointerType(t)
	Printfln("Pointer Type: %v", pt)
	Printfln("Follow pointer type: %v", followPointerType(pt))

	transformString(&name)
	Printfln("Follow pointer value: %v", name)

}
