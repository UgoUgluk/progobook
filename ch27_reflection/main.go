package main

import (
	"reflect"
	//"strings"
	// "fmt"
)

func isInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return true
	}
	return false
}
func isFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func convert(src, target interface{}) (result interface{},
	assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		if (isInt(targetVal) && isInt(srcVal)) &&
			targetVal.OverflowInt(srcVal.Int()) {
			Printfln("Int overflow")
			return src, false
		} else if isFloat(targetVal) && isFloat(srcVal) &&
			targetVal.OverflowFloat(srcVal.Float()) {
			Printfln("Float overflow")
			return src, false
		}
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}
	return
}
func swap(first interface{}, second interface{}) {
	firstValue, secondValue := reflect.ValueOf(first),
		reflect.ValueOf(second)
	if firstValue.Type() == secondValue.Type() &&
		firstValue.Kind() == reflect.Ptr &&
		firstValue.Elem().CanSet() &&
		secondValue.Elem().CanSet() {
		temp := reflect.New(firstValue.Elem().Type())
		temp.Elem().Set(firstValue.Elem())
		firstValue.Elem().Set(secondValue.Elem())
		secondValue.Elem().Set(temp.Elem())
	}
}

func main() {
	name := "Alice"
	price := 279
	city := "London"
	newVal, ok := convert(price, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
	newVal, ok = convert(name, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)

	newVal, ok = convert(5000, int8(100))
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)

	swap(&name, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

}
