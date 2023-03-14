package main

import "fmt"

//ChannelMessage channel for errors
type ChannelMessage struct {
	Category      string
	Total         float64
	CategoryError error
}

// TotalPrice calc sum prices of group
func (slice ProductSlice) TotalPrice(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		//err = errors.New("Cannot find category")
		err = fmt.Errorf("Cannot find category: %v", category)
	}
	return
}

//TotalPriceAsync method for send async data about total price
func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
