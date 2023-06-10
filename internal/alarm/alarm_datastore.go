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

// CreateAlarm creates a new alarm.
func CreateAlarm(userID string) *AlarmEvent {
	return &AlarmEvent{
		ID:          generateID(),
		TriggeredAt: time.Now(),
		Status:      "Triggered",
	}
}

// UpdateStatus updates the status of the alarm.
func (a *AlarmEvent) UpdateStatus(status string) {
	a.Status = status
}

// RetrieveAlarm retrieves the alarm by its ID.
func RetrieveAlarm(id string) *AlarmEvent {
	// Retrieve the alarm from the database or any other storage mechanism
	// using the provided ID.
	// Implement your retrieval logic here.
	return &AlarmEvent{
		ID:          id,
		TriggeredAt: time.Now(),
		Status:      "Triggered",
	}
}

// CancelAlarm cancels an existing alarm.
func (a *AlarmEvent) CancelAlarm() {
	a.Status = "Cancelled"
}

// AcknowledgeAlarm acknowledges the alarm, marking it as acknowledged.
func (a *AlarmEvent) AcknowledgeAlarm() {
	a.Status = "Acknowledged"
}

// ResolveAlarm resolves the alarm, marking it as resolved.
func (a *AlarmEvent) ResolveAlarm() {
	a.Status = "Resolved"
}

// GetAlarmStatus returns the current status of the alarm.
func (a *AlarmEvent) GetAlarmStatus() string {
	return a.Status
}

// SetTriggeredAt sets the triggered timestamp of the alarm.
func (a *AlarmEvent) SetTriggeredAt(timestamp time.Time) {
	a.TriggeredAt = timestamp
}

// GetAllAlarms retrieves all the alarms associated with a user.
func GetAllAlarms(userID string) []*AlarmEvent {
	// Retrieve all the alarms associated with the given userID.
	// Implement your retrieval logic here.
	// Return a slice of Alarm pointers.
	return []*AlarmEvent{
		{ID: "alarm1", TriggeredAt: time.Now(), Status: "Triggered"},
		{ID: "alarm2", TriggeredAt: time.Now(), Status: "Acknowledged"},
		{ID: "alarm3", TriggeredAt: time.Now(), Status: "Resolved"},
	}
}

// generateID generates a unique ID for the alarm.
func generateID() string {
	// Implement your ID generation logic here.
	// This is just a placeholder implementation.
	return "alarmID123"
}
