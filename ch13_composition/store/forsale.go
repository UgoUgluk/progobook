package store

//ItemForSale interface for  product for sale
type ItemForSale interface {
	Price(taxRate float64) float64
}
