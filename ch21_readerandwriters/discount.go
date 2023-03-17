package main

import (
	"encoding/json"
	"strconv"
)

// DiscountedProduct Product with Discount
type DiscountedProduct struct {
	//*Product
	//*Product `json:"product,omitempty"`
	*Product `json:",omitempty"`
	//Discount float64 `json:"-"`
	//Discount float64 `json:",string"`
	Discount float64 `json:"offer,string"`
}

// MarshalJSON other json out
func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}

// UnmarshalJSON other json decode
func (dp *DiscountedProduct) UnmarshalJSON(data []byte) (err error) {
	mdata := map[string]interface{}{}
	err = json.Unmarshal(data, &mdata)
	if dp.Product == nil {
		dp.Product = &Product{}
	}
	if err == nil {
		if name, ok := mdata["Name"].(string); ok {
			dp.Name = name
		}
		if category, ok := mdata["Category"].(string); ok {
			dp.Category = category
		}
		if price, ok := mdata["Price"].(float64); ok {
			dp.Price = price
		}
		if discount, ok := mdata["Offer"].(string); ok {
			fpval, fperr := strconv.ParseFloat(discount, 64)
			if fperr == nil {
				dp.Discount = fpval
			}
		}
	}
	return
}
