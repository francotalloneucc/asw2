package dtos

type HotelDTO struct {
	Name      string   `json:"name" binding:"required"`
	Location  string   `json:"location" binding:"required"`
	Amenities []string `json:"amenities"`
	Rating    float64  `json:"rating"`
}
