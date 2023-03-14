package fmt

import "strconv"

//ToCurrency format float to string with $
func ToCurrency(amount float64) string {
	return "$" + strconv.FormatFloat(amount, 'f', 2, 64)
}
