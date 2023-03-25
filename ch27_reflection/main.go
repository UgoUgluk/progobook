package main

import (
	"reflect"
	"strings"
)

func selectValue(data interface{}, index int) (result interface{}) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}
	return
}
func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(
					elemValue.String()))
			}
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(),
				elemValue)
		}
	}
}

func setAll(src interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() == reflect.Ptr &&
			targetVal.Elem().Type() == srcVal.Type() &&
			targetVal.Elem().CanSet() {
			targetVal.Elem().Set(srcVal)
		}
	}
}

func contains(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
				found = true
			}
		}
	}

	return
}

func main() {
	//names := []string{"Alice", "Bob", "Charlie"}
	//val := selectValue(names, 1).(string)
	//Printfln("Selected: %v", val)
	name := "Alice"
	price := 279
	city := "London"
	incrementOrUpper(name, price, city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

	setAll("New String", &name, &price, &city)
	setAll(10, &name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

	citiesSlice := []string{"Paris", "Rome", "London"}
	Printfln("Found #1: %v", contains(citiesSlice, city))
	sliceOfSlices := [][]string{citiesSlice, {"First", "Second", "Third"}}
	Printfln("Found #2: %v", contains(sliceOfSlices, citiesSlice))

}
