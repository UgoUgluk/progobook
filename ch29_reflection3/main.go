package main

import (
	"reflect"
)

func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType([]int{}, structType)
		}
	}
}
func inspectStructType(baseIndex []int, structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		fieldIndex := append(baseIndex, i)
		field := structType.Field(i)
		Printfln("Field %v: Name: %v, Type: %v, Exported: %v",
			fieldIndex, field.Name, field.Type, field.PkgPath == "")
		if field.Type.Kind() == reflect.Struct {
			field := structType.FieldByIndex(fieldIndex)
			inspectStructType(fieldIndex, field.Type)
		}

	}
	Printfln("--- End Struct Type: %v", structType)
}
func describeField(s interface{}, fieldName string) {
	structType := reflect.TypeOf(s)
	field, found := structType.FieldByName(fieldName)
	if found {
		Printfln("Found: %v, Type: %v, Index: %v",
			field.Name, field.Type, field.Index)
		index := field.Index
		for len(index) > 1 {
			index = index[0 : len(index)-1]
			field = structType.FieldByIndex(index)
			Printfln("Parent : %v, Type: %v, Index: %v",
				field.Name, field.Type, field.Index)
		}
		Printfln("Top-Level Type: %v", structType)
	} else {
		Printfln("Field %v not found", fieldName)
	}
}

func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v: %v", field.Name, tagName,
			valGet)
		Printfln("Field: %v, Tag %v: %v, Set: %v",
			field.Name, tagName, valLookup, ok)
	}
}

// Person present person
type Person struct {
	Name    string `alias:"id"`
	City    string `alias:""`
	Country string
}

func getFieldValues(s interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Struct {
		for i := 0; i < structValue.NumField(); i++ {
			fieldType := structValue.Type().Field(i)
			fieldVal := structValue.Field(i)
			Printfln("Name: %v, Type: %v, Value: %v",
				fieldType.Name, fieldType.Type, fieldVal)
		}
	} else {
		Printfln("Not a struct")
	}
}

func setFieldValue(s interface{}, newVals map[string]interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Ptr &&
		structValue.Elem().Kind() == reflect.Struct {
		for name, newValue := range newVals {
			fieldVal := structValue.Elem().FieldByName(name)
			if fieldVal.CanSet() {
				fieldVal.Set(reflect.ValueOf(newValue))
			} else if fieldVal.CanAddr() {
				ptr := fieldVal.Addr()
				if ptr.CanSet() {
					ptr.Set(reflect.ValueOf(newValue))
				} else {
					Printfln("Cannot set field via pointer")
				}
			} else {
				Printfln("Cannot set field")
			}
		}
	} else {
		Printfln("Not a pointer to a struct")
	}
}

func main() {
	inspectStructs(Purchase{})

	describeField(Purchase{}, "Price")

	inspectTags(Person{}, "alias")

	stringType := reflect.TypeOf("this is a string")
	structType := reflect.StructOf([]reflect.StructField{
		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
		{Name: "City", Type: stringType, Tag: `alias:""`},
		{Name: "Country", Type: stringType},
	})
	inspectTags(reflect.New(structType), "alias")

	product := Product{Name: "Kayak", Category: "Watersports",
		Price: 279}
	customer := Customer{Name: "Acme", City: "Chicago"}
	purchase := Purchase{
		Customer: customer,
		Product:  product,
		Total:    279,
		taxRate:  10}
	setFieldValue(&purchase, map[string]interface{}{
		"City": "London", "Category": "Boats", "Total": 100.50,
	})

	getFieldValues(purchase)

}
