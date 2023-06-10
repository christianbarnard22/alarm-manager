package alarm

import (
	"context"

	"github.com/christianbarnard22/alarm-manager/internal/logging"
	"github.com/christianbarnard22/alarm-manager/internal/noonlight"
	validator "github.com/christianbarnard22/alarm-manager/internal/tools"
)

// AlarmService represents the alarm service.
type Alarm interface {
	CreateAlarm(ctx context.Context, req *CreateAlarmParams) (*AlarmEvent, error)
	GetAlarm(ctx context.Context, id string) *AlarmEvent
	CancelAlarm(ctx context.Context, alarmId string) (*noonlight.CancelAlarmResponse, error)
}

// alarmService is an implementation of the AlarmService interface.
type AlarmService struct {
	NoonlightClient *noonlight.NoonlightClient
}

// NewAlarmService creates a new instance of AlarmService.
func NewAlarmService(noonlightClient *noonlight.NoonlightClient) Alarm {
	return &AlarmService{
		NoonlightClient: noonlightClient,
	}
}

type CreateAlarmParams struct {
	Name     string       `json:"name" validate:"required"`
	Phone    string       `json:"phone" validate:"required"`
	Location LocationData `json:"location" validate:"required"`
}

// CreateAlarm creates a new alarm.
func (s *AlarmService) CreateAlarm(ctx context.Context, req *CreateAlarmParams) (*AlarmEvent, error) {
	// make sure to provide name, Phone number, location object with address
	logger := logging.GetLogger(ctx)

	logger.Info("Creating alarm")

	err := validator.ValidateStruct(req)
	if err != nil {
		logger.Warn("Warning: Alarm creation failed!", err)
		return nil, err
	}

	noonlightReq := &noonlight.TriggerAlarmRequestBody{
		Name:  req.Name,
		Phone: req.Phone,
		Location: &noonlight.LocationData{
			Address: (*noonlight.AddressData)(req.Location.Address),
		},
	}

	resp, err := s.NoonlightClient.TriggerAlarm(ctx, noonlightReq)

	if err != nil {
		logger.Error("Failed to Trigger Alarm")
		return nil, err
	}

	alarmevent := &AlarmEvent{
		ID:          resp.ID,
		Status:      resp.Status,
		TriggeredAt: resp.CreatedAt,
	}

	logger.Info("Alarm successfully trggered")
	return alarmevent, nil
}

// RetrieveAlarm retrieves an alarm by its ID.
func (s *AlarmService) GetAlarm(ctx context.Context, id string) *AlarmEvent {
	return nil
}

func (s *AlarmService) CancelAlarm(ctx context.Context, alarmId string) (*noonlight.CancelAlarmResponse, error) {
	return nil, nil
}
