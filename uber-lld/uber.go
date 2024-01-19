package main

func main() {
	nidheyRider := &Rider{
		name:   "Nidhey",
		rating: "2",
	}
	niluRider := &Rider{
		name:   "Nilu",
		rating: "5",
	}
	riderManagerInstance, _ = GetRiderManager()
	riderManagerInstance.AddRider(nidheyRider.name, nidheyRider)
	riderManagerInstance.AddRider(niluRider.name, niluRider)

	sarthakDriver := &Driver{
		name:   "Sarthak",
		rating: "1",
	}
	shubhnagiDriver := &Driver{
		name:   "Sarthak",
		rating: "1",
	}
	driverManagerInstance, _ := GetDriverManager()
	driverManagerInstance.AddDriver(sarthakDriver.name, sarthakDriver)
	driverManagerInstance.AddDriver(shubhnagiDriver.name, shubhnagiDriver)

	tripManagerInstance, _ := GetTripManager()

	tripManagerInstance.CreateTrip(nidheyRider, "Vaishali Nagar", "Sadar")
	tripManagerInstance.CreateTrip(niluRider, "Manish Nagar", "Hingna")

	tripsMap := tripManagerInstance.GetTripsMap()
	for _, trip := range tripsMap {
		trip.DisplayTripDetails()
	}
}
