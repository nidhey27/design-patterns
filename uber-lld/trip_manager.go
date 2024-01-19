package main

import (
	"fmt"
	"sync"
)

var (
	tripManagerInstance *TripManager
)

type TripManager struct {
	mtx               sync.Mutex
	RideManager       *RiderManager
	DriverManager     *DriverManager
	TripsMetaDataInfo map[int]*TripMetaData
	TripsInfo         map[int]*Trip
}

func GetTripManager() (*TripManager, error) {
	err := fmt.Errorf("")
	if tripManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if tripManagerInstance == nil {
			tripManagerInstance = &TripManager{
				TripsMetaDataInfo: make(map[int]*TripMetaData),
				TripsInfo: make(map[int]*Trip),
			}
		} else {
			err = fmt.Errorf("TripManager instance already created.")
		}
	} else {
		err = fmt.Errorf("TripManager instance already created.")
	}
	return tripManagerInstance, err
}

func (tm *TripManager) CreateTrip(rider *Rider, src string, dst string) {
	metaData := &TripMetaData{
		SrcLocation: src,
		DstLocation: dst,
		RiderRating: rider.GetRating(),
	}

	strategyManager, _ := GetStrategyManager()
	pricingStrategy := strategyManager.DeterminePricingStrategy(metaData)
	driverMatchingStrategy := strategyManager.DetermineMatchingStrategy(metaData)

	driver := driverMatchingStrategy.MatchDriver(metaData)
	tripPrice := pricingStrategy.CalculatePrice(metaData)

	trip := &Trip{
		TripID: len(tm.TripsInfo) + 1,
		Rider:                  rider,
		Driver:                 driver,
		SrcLocation:            src,
		DstLocation:            dst,
		Price:                  tripPrice,
		PricingStrategy:        &pricingStrategy,
		DriverMatchingStrategy: &driverMatchingStrategy,
	}
	tripId := trip.GetTripID()
	tm.TripsInfo[tripId] = trip
	tm.TripsMetaDataInfo[tripId] = metaData
}

func (tm *TripManager) GetTripsMap() map[int]*Trip {
	return tm.TripsInfo
}
