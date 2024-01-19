package main

type Driver struct {
	name      string
	rating    string
	available bool
}

func (d *Driver) UpdateAvailable(status bool) {
	d.available = status
}

func (d *Driver) GetDriverName() string {
	return d.name
}

func (d *Driver) GetRating() string {
	return d.rating
}
