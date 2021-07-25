package danalto

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// TODO: Make sure think about what should be private vs public
// TODO: What about the optional start and end time
// TODO: What about versioning?
// TODO: Do I need all the headers in sendRequest?


// TODO: Are these two types actually needed?
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

const (
	baseURL = "https://three.danalto.com/api/v1"
	apiKey = "U2VhbXVzQm9ubmVyOmFycmFubW9yZUlvVA=="
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}


// TODO: Changed from this (note the pointer - *SensorDataList)
// TODO: Do I need to pass the interface in to solve this issue or just return a copy as below
// TODO: func (c *Client) GetDeviceData(apiKey string, deviceId string/*, opts *DeviceDataOpts*/) (*SensorDataList, error) {

func (c *Client) GetDeviceData(apiKey string, deviceId string/*, opts *DeviceDataOpts*/) (SensorDataList, error) {

	// TODO: Need to add start / end timestamp as query paramseters
	// https://three.danalto.com/api/v1/devices/a81758fffe0346cd/data/?start_timestamp=1626821501651&end_timestamp=1626821589651

	// req, err := http.NewRequest("GET", fmt.Sprintf("%s/faces?limit=%d&page=%d", baseURL), nil)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/devices/%s/data/", baseURL, deviceId), nil)
	if err != nil {
		return nil, err
	}

	res := SensorDataList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	// TODO: Should I use a pointer reference here or not
	// TODO: return &res, nil
	return res, nil
}

// TODO: Add some logging here.
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	// TODO: Pass in apiKey / username&password soemehow
	// req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.apiKey))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", apiKey))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	//fullResponse := successResponse{
	//	Data: v,
	//
	//if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
	//	return err
	//}

	// TODO: In debug it is saying that 'v is shadowed'. Find out why.
	// Removed the pointer dereference (&v)

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

// TODO: Should this be in a separate file somewhere?
// TODO: Do I need username/password on every request?
// TODO: Where can it be stored if don't want it on every request?

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
