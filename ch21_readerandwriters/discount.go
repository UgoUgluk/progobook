package main

import "encoding/json"

//DiscountedProduct Product with Discount
type DiscountedProduct struct {
	//*Product
	//*Product `json:"product,omitempty"`
	*Product `json:",omitempty"`
	//Discount float64 `json:"-"`
	Discount float64 `json:",string"`
}

//MarshalJSON other json out
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
