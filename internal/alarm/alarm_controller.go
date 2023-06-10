package alarm

import (
	"context"
	"fmt"
)

// AlarmHandler is responsible for handling alarm-related requests.
type AlarmHandler struct {
	alarmService AlarmService
}

// NewAlarmHandler creates a new instance of AlarmHandler.
func NewAlarmHandler(alarmService AlarmService) *AlarmHandler {
	return &AlarmHandler{
		alarmService: alarmService,
	}
}

// CreateAlarmResponse is the response structure for creating a new alarm.
type CreateAlarmResponse struct {
	AlarmID string
}

// CreateAlarm creates a new alarm based on the provided request.
func (h *AlarmHandler) CreateAlarm(ctx context.Context, req *CreateAlarmParams) (*CreateAlarmResponse, error) {
	// logger := logging.GetLogger(ctx)

	alarm, err := h.alarmService.CreateAlarm(ctx, req)
	if err != nil {
		return nil, err
	}
	// Store the alarm in the database or any other storage mechanism.
	// Implement your storage logic here.
	// ...

	return &CreateAlarmResponse{
		AlarmID: alarm.ID,
	}, nil
}

// GetAlarmRequest is the request structure for retrieving an alarm.
type GetAlarmRequest struct {
	AlarmID string
}

// GetAlarmResponse is the response structure for retrieving an alarm.
type GetAlarmResponse struct {
	AlarmEvent *AlarmEvent
}

// GetAlarm retrieves the alarm based on the provided request.
func (h *AlarmHandler) GetAlarm(ctx context.Context, req *GetAlarmRequest) (*GetAlarmResponse, error) {
	alarm := h.alarmService.GetAlarmStatus(ctx, req.AlarmID)
	if alarm == nil {
		return nil, fmt.Errorf("alarm not found")
	}

	return &GetAlarmResponse{
		AlarmEvent: alarm,
	}, nil
}

// CancelAlarmRequest is the request structure for canceling an alarm.
type CancelAlarmRequest struct {
	AlarmID string
}

// CancelAlarm cancels the alarm based on the provided request.
func (h *AlarmHandler) CancelAlarm(ctx context.Context, req *CancelAlarmRequest) error {
	// Retrieve the alarm from the database or any other storage mechanism.
	// Implement your retrieval logic here.
	// alarm, err := h.alarmService.CancelAlarm(ctx, req.AlarmID)
	// if err != nil {
	// 	return .Error("Unable to cancel Alarm: ", err)
	// }

	// // Cancel the alarm.
	// alarm.CancelAlarm()

	// Store the updated alarm in the database or any other storage mechanism.
	// Implement your storage logic here.
	// ...

	return nil
}
