package main

import (
	"fmt"
	"reflect"
	"strings"
)

func inspectFuncType(f interface{}) {
	funcType := reflect.TypeOf(f)
	if funcType.Kind() == reflect.Func {
		Printfln("Function parameters: %v", funcType.NumIn())
		for i := 0; i < funcType.NumIn(); i++ {
			paramType := funcType.In(i)
			if i < funcType.NumIn()-1 {
				Printfln("Parameter #%v, Type: %v", i,
					paramType)
			} else {
				Printfln("Parameter #%v, Type: %v, Variadic: %v", i, paramType,
					funcType.IsVariadic())
			}
		}
		Printfln("Function results: %v", funcType.NumOut())
		for i := 0; i < funcType.NumOut(); i++ {
			resultType := funcType.Out(i)
			Printfln("Result #%v, Type: %v", i, resultType)
		}
	}
}
func invokeFunction(f interface{}, params ...interface{}) {
	paramVals := []reflect.Value{}
	for _, p := range params {
		paramVals = append(paramVals, reflect.ValueOf(p))
	}
	funcVal := reflect.ValueOf(f)
	if funcVal.Kind() == reflect.Func {
		results := funcVal.Call(paramVals)
		for i, r := range results {
			Printfln("Result #%v: %v", i, r)
		}
	}
}

func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
	sliceVal := reflect.ValueOf(slice)
	mapperVal := reflect.ValueOf(mapper)
	mapped = []interface{}{}
	if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func {
		paramTypes := []reflect.Type{sliceVal.Type().Elem()}
		resultTypes := []reflect.Type{}
		for i := 0; i < mapperVal.Type().NumOut(); i++ {
			resultTypes = append(resultTypes, mapperVal.Type().Out(i))
		}
		expectedFuncType := reflect.FuncOf(paramTypes, resultTypes, mapperVal.Type().IsVariadic())
		if mapperVal.Type() == expectedFuncType {
			for i := 0; i < sliceVal.Len(); i++ {
				result := mapperVal.Call([]reflect.Value{sliceVal.Index(i)})
				for _, r := range result {
					mapped = append(mapped, r.Interface())
				}
			}
		} else {
			Printfln("Function type not as expected")
		}
	}
	return
}

func makeMapperFunc(mapper interface{}) interface{} {
	mapVal := reflect.ValueOf(mapper)
	if mapVal.Kind() == reflect.Func &&
		mapVal.Type().NumIn() == 1 &&
		mapVal.Type().NumOut() == 1 {

		inType := reflect.SliceOf(mapVal.Type().In(0))
		inTypeSlice := []reflect.Type{inType}
		outType := reflect.SliceOf(mapVal.Type().Out(0))
		outTypeSlice := []reflect.Type{outType}
		funcType := reflect.FuncOf(inTypeSlice, outTypeSlice, false)
		funcVal := reflect.MakeFunc(funcType,
			func(params []reflect.Value) (results []reflect.Value) {
				srcSliceVal := params[0]
				resultsSliceVal := reflect.MakeSlice(outType, srcSliceVal.Len(), 10)
				for i := 0; i < srcSliceVal.Len(); i++ {
					r := mapVal.Call([]reflect.Value{srcSliceVal.Index(i)})
					resultsSliceVal.Index(i).Set(r[0])
				}
				results = []reflect.Value{resultsSliceVal}
				return
			})
		return funcVal.Interface()
	}
	Printfln("Unexpected types")
	return nil
}

func main() {
	inspectFuncType(Find)

	names1 := []string{"Alice", "Bob", "Charlie"}
	invokeFunction(Find, names1, "London", "Bob")

	results1 := mapSlice(names1, strings.ToUpper)
	Printfln("Results: %v", results1)

	lowerStringMapper := makeMapperFunc(strings.ToLower).(func([]string) []string)
	names := []string{"Alice", "Bob", "Charlie"}
	results := lowerStringMapper(names)
	Printfln("Lowercase Results: %v", results)
	incrementFloatMapper := makeMapperFunc(func(val float64) float64 {
		return val + 1
	}).(func([]float64) []float64)
	prices := []float64{279, 48.95, 19.50}
	floatResults := incrementFloatMapper(prices)
	Printfln("Increment Results: %v", floatResults)
	floatToStringMapper := makeMapperFunc(func(val float64) string {
		return fmt.Sprintf("$%.2f", val)
	}).(func([]float64) []string)
	Printfln("Price Results: %v", floatToStringMapper(prices))

}
