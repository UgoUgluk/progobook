package main

import (
	"reflect"
)

func checkImplementation(check interface{}, targets ...interface{}) {
	checkType := reflect.TypeOf(check)
	if checkType.Kind() == reflect.Ptr &&
		checkType.Elem().Kind() == reflect.Interface {
		checkType := checkType.Elem()
		for _, target := range targets {
			targetType := reflect.TypeOf(target)
			Printfln("Type %v implements %v: %v",
				targetType, checkType, targetType.Implements(checkType))
		}
	}
}

type wrapper struct {
	NamedItem
}

func getUnderlying(item wrapper, fieldName string) {
	itemVal := reflect.ValueOf(item)
	fieldVal := itemVal.FieldByName(fieldName)
	Printfln("Field Type: %v", fieldVal.Type())
	for i := 0; i < fieldVal.Type().NumMethod(); i++ {
		method := fieldVal.Type().Method(i)
		Printfln("Interface Method: %v, Exported: %v",
			method.Name, method.PkgPath == "")
	}
	Printfln("--------")

	if fieldVal.Kind() == reflect.Interface {
		Printfln("Underlying Type: %v",
			fieldVal.Elem().Type())
		for i := 0; i < fieldVal.Elem().Type().NumMethod(); i++ {
			method := fieldVal.Elem().Type().Method(i)
			Printfln("Underlying Method: %v", method.Name)
		}
	}
}

func main() {
	currencyItemType := (*CurrencyItem)(nil)
	checkImplementation(currencyItemType, Product{}, &Product{}, &Purchase{})

	getUnderlying(wrapper{NamedItem: &Product{}}, "NamedItem")

}
