package main

import (
	"encoding/json"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}

	//num
	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}

	//with types
	reader2 := strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var ival int
	vals2 := []interface{}{&bval, &sval, &fpval, &ival}
	decoder2 := json.NewDecoder(reader2)
	for i := 0; i < len(vals2); i++ {
		err := decoder2.Decode(vals2[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)

	//arrays
	reader3 := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	vals3 := []interface{}{}
	decoder3 := json.NewDecoder(reader3)
	for {
		var decodedVal interface{}
		err := decoder3.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals3 = append(vals3, decodedVal)
	}
	for _, val := range vals3 {
		Printfln("Decoded (%T): %v", val, val)
	}

	//maps
	reader4 := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	//m := map[string]interface{}{}
	m := map[string]float64{}
	decoder4 := json.NewDecoder(reader4)
	err := decoder4.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}

	//struct
	reader5 := strings.NewReader(`
		{"Name":"Kayak","Category":"Watersports","Price":279}
		{"Name":"Lifejacket","Category":"Watersports" }
		{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
	`)
	decoder5 := json.NewDecoder(reader5)
	decoder5.DisallowUnknownFields()
	for {
		var val Product
		err := decoder5.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		}
	}

	//struct with tags
	reader6 := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer":"10"}`)
	decoder6 := json.NewDecoder(reader6)
	for {
		var val DiscountedProduct
		err := decoder6.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v, Discount: %v",
				val.Name, val.Category, val.Price, val.Discount)
		}
	}

}
