package main

type Rider struct {
	name   string
	rating string
}

func (r *Rider) GetRiderName() string {
	return r.name
}

func (r *Rider) GetRating() string {
	return r.rating
}
