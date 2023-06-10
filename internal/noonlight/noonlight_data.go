package noonlight

import (
	"context"
	"net/http"
	"time"
)

// NoonlightAPI represents the client for interacting with the Noonlight API.
type NoonlightAPI interface {
	TriggerAlarm(ctx context.Context, address string) (*TriggerAlarmResponse, error)
	CancelAlarm(ctx context.Context, alarmID string) (*CancelAlarmResponse, error)
	GetAlarmStatus(ctx context.Context, alarmID string) (*GetAlarmStatusResponse, error)
}

type TriggerAlarmRequestBody struct {
	Name     string       `json:"name"`
	Phone    string       `json:"phone"`
	Location LocationData `json:"location"`
}

type LocationData struct {
	Coordinates *CoordinatesData `json:"coordinates,omitempty"`
	Address     *AddressData     `json:"address,omitempty"`
}

type CoordinatesData struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Accuracy int     `json:"accuracy"`
}

type AddressData struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	City  string `json:"city"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}

type TriggerAlarmResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	OwnerID   string    `json:"owner_id"`
}

type NoonlightClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}
