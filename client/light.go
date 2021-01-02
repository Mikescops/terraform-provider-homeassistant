package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Light struct
type Light struct {
	ID         string          `json:"entity_id,omitempty"`
	State      string          `json:"state,omitempty"`
	Attributes LightAttributes `json:"attributes,omitempty"`
}

// LightAttributes struct
type LightAttributes struct {
	Brightness int   `json:"brightness,omitempty"`
	RgbColor   []int `json:"rgb_color,omitempty"`
}

// GetLightState get state of a light device
func (c *Client) GetLightState(lightID string) (*Light, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/states/%s", c.HostURL, lightID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	light := Light{}
	err = json.Unmarshal(body, &light)
	if err != nil {
		return nil, err
	}

	return &light, nil
}

// LightParams params of SetLightState
type LightParams struct {
	ID         string        `json:"entity_id,omitempty"`
	Brightness int           `json:"brightness,omitempty"`
	RgbColor   []interface{} `json:"rgb_color,omitempty"`
}

// SetLightState set state of a light device
func (c *Client) SetLightState(lightParams LightParams, stateParam string) ([]Light, error) {

	var state string = "off"
	rb, err := json.Marshal(LightParams{
		ID: lightParams.ID,
	})

	if stateParam == "on" {
		state = "on"
		rb, err = json.Marshal(lightParams)
	}
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/light/turn_%s", c.HostURL, state), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	light := []Light{}
	err = json.Unmarshal(body, &light)
	if err != nil {
		return nil, err
	}

	return light, nil
}
