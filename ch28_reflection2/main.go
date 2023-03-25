package main

import (
	"reflect"
)

func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)
	if arrayOrSliceVal.Kind() == reflect.Slice {
		elemVal := arrayOrSliceVal.Index(index)
		if elemVal.CanSet() {
			elemVal.Set(replacementVal)
		}
	} else if arrayOrSliceVal.Kind() == reflect.Ptr &&
		arrayOrSliceVal.Elem().Kind() == reflect.Array &&
		arrayOrSliceVal.Elem().CanSet() {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}

func checkElemType(val interface{}, arrOrSlice interface{}) bool {
	elemType := reflect.TypeOf(val)
	arrOrSliceType := reflect.TypeOf(arrOrSlice)
	return (arrOrSliceType.Kind() == reflect.Array ||
		arrOrSliceType.Kind() == reflect.Slice) &&
		arrOrSliceType.Elem() == elemType
}

func enumerateStrings(arrayOrSlice interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	if (arrayOrSliceVal.Kind() == reflect.Array ||
		arrayOrSliceVal.Kind() == reflect.Slice) &&
		arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
		for i := 0; i < arrayOrSliceVal.Len(); i++ {
			Printfln("Element: %v, Value: %v", i,
				arrayOrSliceVal.Index(i).String())
		}
	}
}

func findAndSplit(slice interface{}, target interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				return sliceVal.Slice(0, i+1)
			}
		}
	}
	return slice
}

func pickValues(slice interface{}, indices ...int) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
		for _, index := range indices {
			newSlice = reflect.Append(newSlice,
				sliceVal.Index(index))
		}
		return newSlice
	}
	return nil
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}
	Printfln("Slice (string): %v", checkElemType("testString", slice))
	Printfln("Array (string): %v", checkElemType("testString", array))
	Printfln("Array (int): %v", checkElemType(10, array))

	Printfln("Original slice: %v", slice)
	newCity := "Paris"
	setValue(slice, 1, newCity)
	Printfln("Modified slice: %v", slice)
	Printfln("Original slice: %v", array)
	newCity = "Rome"
	setValue(&array, 1, newCity)
	Printfln("Modified slice: %v", array)

	enumerateStrings(slice)
	enumerateStrings(array)

	Printfln("Strings: %v", findAndSplit(slice, "London"))
	numbers := []int{1, 3, 4, 5, 7}
	Printfln("Numbers: %v", findAndSplit(numbers, 4))

	slice2 := []string{name, city, hobby, "Bob", "Paris", "Soccer"}
	picked := pickValues(slice2, 0, 3, 5)
	Printfln("Picked values: %v", picked)

}
