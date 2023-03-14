package main

//CategoryError error handling
type CategoryError struct {
	requestedCategory string
}

func (e *CategoryError) Error() string {
	return "Category " + e.requestedCategory + " does not exist"
}

//ChannelMessage channel for errors
type ChannelMessage struct {
	Category string
	Total    float64
	*CategoryError
}

// TotalPrice calc sum prices of group
func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = &CategoryError{requestedCategory: category}
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
