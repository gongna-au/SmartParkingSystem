package requests

type ParkingSearchByLatitudeAndLongitude struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
} //@name ParkingSearchByLatitudeAndLongitude
