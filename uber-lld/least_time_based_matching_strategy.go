package main

import "fmt"

type LeastTimeBasedMatchingStrategy struct{}

func (l *LeastTimeBasedMatchingStrategy) MatchDriver(metadata *TripMetaData) *Driver {

	driverManager, _ := GetDriverManager()


	if len(driverManager.GetDriversMap()) == 0 {
		fmt.Println("No drivers! What service is this huh?")
		return nil
	}

	fmt.Println("Using quadtree to see nearest cabs, using driver manager to get details of drivers and send notifications")
	driverMap := driverManager.GetDriversMap()
	var keys []string
	for key := range driverMap {
		keys = append(keys, key)
	}
	if len(keys) > 0 {
		// Get the first key from the slice
		firstKey := keys[0]

		// Access the corresponding value from the map
		driver := driverMap[firstKey]
		fmt.Println("Setting", driver.GetDriverName(), "as driver")
		metadata.SetDriverRating(driver.GetRating())
		return driver
	} else {
		fmt.Println("Map is empty")
		return nil
	}

	return nil
}
