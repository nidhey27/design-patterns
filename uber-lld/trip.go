package main

import "fmt"

type Trip struct {
	Rider                  *Rider
	Driver                 *Driver
	SrcLocation            string
	DstLocation            string
	Status                 bool
	TripID                 int
	Price                  float64
	PricingStrategy        *PricingStrategy
	DriverMatchingStrategy *DriverMatchingStrategy
}

func (t *Trip) GetTripID() int {
	return t.TripID
}

func (t *Trip) DisplayTripDetails() {
	fmt.Println("-------------- TRIP DETAILS -----------")
	fmt.Println("Trip ID:", t.TripID)
	fmt.Println("Rider:", t.Rider.GetRiderName())
	fmt.Println("Driver:", t.Driver.GetDriverName())
	fmt.Println("Source Location:", t.SrcLocation)
	fmt.Println("Destination Location:", t.DstLocation)
	fmt.Println("Status:", t.Status)
	fmt.Println("Price:", t.Price)
	fmt.Println("---------------------------------------")
}
