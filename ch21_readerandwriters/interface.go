package main

//Named interface for name
type Named interface{ GetName() string }

//Person interface for PersonName
type Person struct{ PersonName string }

//GetName get person name
func (p *Person) GetName() string { return p.PersonName }

//GetName get discounted product name
func (p *DiscountedProduct) GetName() string { return p.Name }
