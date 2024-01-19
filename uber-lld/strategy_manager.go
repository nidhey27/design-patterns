package main

import (
	"fmt"
	"sync"
)

var (
	strategyManagerInstance *StrategyManager
)

type StrategyManager struct {
	mtx sync.Mutex
}

func GetStrategyManager() (*StrategyManager, error) {
	err := fmt.Errorf("")
	if strategyManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if strategyManagerInstance == nil {
			strategyManagerInstance = &StrategyManager{}
		} else {
			err = fmt.Errorf("StrategyManager instance already created.")
		}
	} else {
		err = fmt.Errorf("StrategyManager instance already created.")
	}
	return strategyManagerInstance, err
}

func (sm *StrategyManager) DeterminePricingStrategy(metaData *TripMetaData) PricingStrategy {
	fmt.Println("Based on location & other factors, setting pricing strategy")
	return &RatingBasedPricingStrategy{}
}

func (sm *StrategyManager) DetermineMatchingStrategy(metaData *TripMetaData) DriverMatchingStrategy {
	fmt.Println("Based on location & other factors, setting driver matching strategy")
	return &LeastTimeBasedMatchingStrategy{}
}
