package main

import "fmt"

type DefaultPricingStrategy struct{}

func (d *DefaultPricingStrategy) CalculatePrice(metadata *TripMetaData) float64 {
	price := 99.99
	fmt.Println("Based on default strategy, price is", price)
	return price
}
