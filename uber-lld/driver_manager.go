package main

import (
	"fmt"
	"sync"
)

var (
	driverManagerInstance *DriverManager
	lock                  = &sync.Mutex{}
)

type DriverManager struct {
	mtx        sync.Mutex
	driversMap map[string]*Driver
}

func GetDriverManager() (*DriverManager, error) {
	err := fmt.Errorf("")
	if driverManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if driverManagerInstance == nil {
			driverManagerInstance = &DriverManager{
				driversMap: make(map[string]*Driver),
			}
		} else {
			err = fmt.Errorf("DriverManager instance already created.")
		}
	} else {
		err = fmt.Errorf("DriverManager instance already created.")
	}
	return driverManagerInstance, err
}

func (rm *DriverManager) AddDriver(driverName string, driver *Driver) {
	rm.mtx.Lock()
	rm.driversMap[driverName] = driver
	rm.mtx.Unlock()
}

func (rm *DriverManager) GetDriver(driverName string) *Driver {
	return rm.driversMap[driverName]
}

func (rm *DriverManager) GetDriversMap() map[string]*Driver {
	return rm.driversMap
}
