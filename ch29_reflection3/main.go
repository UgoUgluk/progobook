package main

import (
	"reflect"
)

func inspectMethods(s interface{}) {
	sType := reflect.TypeOf(s)
	if sType.Kind() == reflect.Struct ||
		(sType.Kind() == reflect.Ptr &&
			sType.Elem().Kind() == reflect.Struct) {
		Printfln("Type: %v, Methods: %v", sType, sType.NumMethod())
		for i := 0; i < sType.NumMethod(); i++ {
			method := sType.Method(i)
			Printfln("Method name: %v, Type: %v", method.Name, method.Type)
		}
	}
}

func executeFirstVoidMethod(s interface{}) {
	sVal := reflect.ValueOf(s)
	for i := 0; i < sVal.NumMethod(); i++ {
		method := sVal.Type().Method(i)
		if method.Type.NumIn() == 1 {
			results := method.Func.Call([]reflect.Value{sVal})
			Printfln("Type: %v, Method: %v, Results: %v", sVal.Type(), method.Name, results)
			break
		} else {
			Printfln("Skipping method %v %v", method.Name, method.Type.NumIn())
		}
	}
}

func main() {
	inspectMethods(Purchase{})
	inspectMethods(&Purchase{})

	executeFirstVoidMethod(&Product{Name: "Kayak", Price: 279})

}
