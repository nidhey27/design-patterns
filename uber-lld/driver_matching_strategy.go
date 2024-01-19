package main

type DriverMatchingStrategy interface {
	MatchDriver(metadata *TripMetaData) *Driver
}
