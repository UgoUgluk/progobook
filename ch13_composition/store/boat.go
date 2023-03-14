package store

//Boat expands product for sale
type Boat struct {
	*Product
	Capacity  int
	Motorized bool
	Name      string
}

//NewBoat create new boat
func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
		"teest",
	}
}
