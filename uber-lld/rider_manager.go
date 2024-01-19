package main

import (
	"fmt"
	"sync"
)

var (
	riderManagerInstance *RiderManager
)

type RiderManager struct {
	mtx       sync.Mutex
	ridersMap map[string]*Rider
}

func GetRiderManager() (*RiderManager, error) {
	err := fmt.Errorf("")
	if riderManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if riderManagerInstance == nil {
			riderManagerInstance = &RiderManager{
				ridersMap: make(map[string]*Rider),
			}
		} else {
			err = fmt.Errorf("RiderManager instance already created.")
		}
	} else {
		err = fmt.Errorf("RiderManager instance already created.")
	}
	return riderManagerInstance, err
}

func (rm *RiderManager) AddRider(riderName string, rider *Rider) {
	rm.mtx.Lock()
	rm.ridersMap[riderName] = rider
	rm.mtx.Unlock()
}

func (rm *RiderManager) GetRider(riderName string) *Rider {
	return rm.ridersMap[riderName]
}
