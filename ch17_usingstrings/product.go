package main

//Product presente a product
type Product struct {
	Name, Category string
	Price          float64
}

//change %v and %+v
/*func (p Product) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}*/

//change %#v
/*func (p Product) GoString() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}*/

//Kayak presente a product Kayak
var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    275,
}

//Products presente a products
var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}
