package main

type TripMetaData struct {
	RiderRating  string
	DriverRating string
	SrcLocation  string
	DstLocation  string
}

func (tmd *TripMetaData) GetRiderRating() string {
	return tmd.RiderRating
}

func (tmd *TripMetaData) GetDriverRating() string {
	return tmd.DriverRating
}

func (tmd *TripMetaData) SetDriverRating(rating string) {
	tmd.DriverRating = rating
}
