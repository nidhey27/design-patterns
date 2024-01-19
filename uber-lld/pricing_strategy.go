package main

type PricingStrategy interface {
	CalculatePrice(metadata *TripMetaData) float64
}
