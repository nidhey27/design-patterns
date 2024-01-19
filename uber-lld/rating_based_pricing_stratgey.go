package main

import (
	"fmt"
	"strconv"
)

type RatingBasedPricingStrategy struct{}

func (r *RatingBasedPricingStrategy) CalculatePrice(metadata *TripMetaData) float64 {
	var price float64 = 0.0

	rating, _ := strconv.Atoi(metadata.RiderRating)

	if rating >= 4 {
		price = 55.0
	} else {
		price = 65.0
	}
	fmt.Println("Based on rating based pricing strategy, price is", price)
	return price
}
