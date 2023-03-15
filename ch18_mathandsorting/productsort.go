package main

import "sort"

//Product discribe a product
type Product struct {
	Name  string
	Price float64
}

//ProductSlice present a product slice
type ProductSlice []Product

//ProductSlices sort a product slice
func ProductSlices(p []Product) {
	sort.Sort(ProductSlice(p))
}

//ProductSlicesAreSorted present a ProductSlicesAreSorted
func ProductSlicesAreSorted(p []Product) {
	sort.IsSorted(ProductSlice(p))
}
func (products ProductSlice) Len() int {
	return len(products)
}
func (products ProductSlice) Less(i, j int) bool {
	return products[i].Price < products[j].Price
}
func (products ProductSlice) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//ProductSliceName ProductSlice by names
type ProductSliceName struct{ ProductSlice }

//ProductSlicesByName ProductSlices by names
func ProductSlicesByName(p []Product) {
	sort.Sort(ProductSliceName{p})
}

//Less for ProductSliceName for sort Interface
func (p ProductSliceName) Less(i, j int) bool {
	return p.ProductSlice[i].Name < p.ProductSlice[j].Name
}

//ProductComparison compare products
type ProductComparison func(p1, p2 Product) bool

//ProductSliceFlex flex struct for ProductSlice and ProductComparison
type ProductSliceFlex struct {
	ProductSlice
	ProductComparison
}

//Less for ProductSliceFlex for sort Interface
func (flex ProductSliceFlex) Less(i, j int) bool {
	return flex.ProductComparison(flex.ProductSlice[i],
		flex.ProductSlice[j])
}

//SortWith sort products for ProductComparison
func SortWith(prods []Product, f ProductComparison) {
	sort.Sort(ProductSliceFlex{prods, f})
}
