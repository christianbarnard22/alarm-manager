package alarm

import "time"

// Alarm represents the alarm entity.
type AlarmEvent struct {
	ID          string
	TriggeredAt time.Time
	Status      string
}

type LocationData struct {
	Coordinates *CoordinatesData `json:"coordinates,omitempty"`
	Address     *AddressData     `json:"address,omitempty"`
}

type CoordinatesData struct {
	Lat      float64 `json:"lat" validate:"required"`
	Lng      float64 `json:"lng" validate:"required"`
	Accuracy int     `json:"accuracy" validate:"required"`
}

type AddressData struct {
	Line1 string `json:"line1" validate:"required"`
	Line2 string `json:"line2"`
	City  string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
	Zip   string `json:"zip" validate:"required"`
}
