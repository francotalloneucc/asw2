package models

type Hotel struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Amenities []string `json:"amenities"`
	Rating   float64  `json:"rating"`
}
