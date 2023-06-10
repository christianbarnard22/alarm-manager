package noonlight

import (
	"context"
	"net/http"
	"time"
)

// NoonlightAPI represents the client for interacting with the Noonlight API.
type NoonlightAPI interface {
	CreateAlarm(ctx context.Context, req CreateAlarmResponse) (*CreateAlarmResponse, error)
	CancelAlarm(ctx context.Context, alarmID string) (string, error)
	GetAlarmStatus(ctx context.Context, alarmID string) (*GetAlarmStatusResponse, error)
}

type CreateAlarmRequestBody struct {
	Name     string        `json:"name" validate:"required"`
	Phone    string        `json:"phone" validate:"required"`
	Location *LocationData `json:"location" validate:"required"`
}

type LocationData struct {
	Coordinates *CoordinatesData `json:"coordinates,omitempty" validate:"required"`
	Address     *AddressData     `json:"address,omitempty" validate:"required"`
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

type CreateAlarmResponse struct {
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
