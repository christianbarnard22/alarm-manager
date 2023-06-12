package noonlight

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/christianbarnard22/alarm-manager/internal/logging"
)

func NewNoonlightClient(apiKey, baseURL string) *NoonlightClient {
	return &NoonlightClient{
		APIKey:  apiKey,
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// CreateAlarm - makes a request to noonlight to trigger an alarm
func (c *NoonlightClient) CreateAlarm(ctx context.Context, requestBody *CreateAlarmRequestBody) (*CreateAlarmResponse, error) {
	logger := logging.GetLogger(ctx)
	// Create the request body JSON
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		logger.Error("Failed to Marshal CreateAlarm request body: ", err)
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/alarms", bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Error("Failed to Create CreateAlarm Request: ", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.APIKey)

	// Send the request using the client's HTTPClient
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		logger.Error("Failed to complete CreateAlarm Request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and process the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to Read CreateAlarm Response: ", err)
		return nil, err
	}

	// Unmarshal the response JSON
	var alarmResponse CreateAlarmResponse
	err = json.Unmarshal(respBody, &alarmResponse)
	if err != nil {
		logger.Error("Failed to unMarshal CancelAlarm Response: ", err)
		return nil, err
	}

	return &alarmResponse, nil
}

type CancelAlarmResponse struct {
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// CancelAlarm - updaets a triggered alar to be canceled.
func (client *NoonlightClient) CancelAlarm(ctx context.Context, alarmID string) (*CancelAlarmResponse, error) {
	logger := logging.GetLogger(ctx)
	url := fmt.Sprintf("%s/dispatch/v1/alarms/%s/status", client.BaseURL, alarmID)

	requestBody := struct {
		Status string `json:"status"`
	}{
		Status: "CANCELED",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		logger.Error("Failed to Marshal CancelAlarm request body: ", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		logger.Error("Failed to create CancelAlarm request: ", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", client.APIKey)

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Failed to complete CancelAlarm Request: ", err)
		return nil, fmt.Errorf("failed to cancel alarm: %s", resp.Status)
	}

	cancelAlarmResponse := &CancelAlarmResponse{}
	err = json.NewDecoder(resp.Body).Decode(cancelAlarmResponse)
	if err != nil {
		return nil, err
	}

	logger.Info("Alarm Canceled")
	return cancelAlarmResponse, nil
}

type GetAlarmStatusResponse struct {
	Status string `json:"status"`
}

// ...

func (c *NoonlightClient) GetAlarmStatus(ctx context.Context, alarmID string) (*GetAlarmStatusResponse, error) {
	logger := logging.GetLogger(ctx)
	// Construct the URL for the GET request
	url := fmt.Sprintf("%s/dispatch/v1/alarms/%s/status", c.BaseURL, alarmID)

	// Make the GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("Failed to create getAlarmStatus request: ", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		logger.Error("Failed complete getAlarmStatus request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body
	var statusResp struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(resp.Body).Decode(&statusResp)
	if err != nil {
		logger.Error("Failed to decode getAlarmStatus request: ", err)
		return nil, err
	}

	return &GetAlarmStatusResponse{
		Status: statusResp.Status,
	}, nil
}
