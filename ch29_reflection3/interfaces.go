package main

import "fmt"

//NamedItem present named item
type NamedItem interface {
	GetName() string
	unexportedMethod()
}

//CurrencyItem present currency item
type CurrencyItem interface {
	GetAmount() string
	currencyName() string
}

//GetName get name of product
func (p *Product) GetName() string {
	return p.Name
}

//GetName get name of customer
func (c *Customer) GetName() string {
	return c.Name
}

//GetAmount get amount of product
func (p *Product) GetAmount() string {
	return fmt.Sprintf("$%.2f", p.Price)
}

func (p *Product) currencyName() string {
	return "USD"
}
func (p *Product) unexportedMethod() {}
