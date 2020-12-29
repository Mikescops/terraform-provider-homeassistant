package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Light struct {
	ID    string `json:"entity_id,omitempty"`
	State string `json:"state,omitempty"`
}

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

type LightParams struct {
	ID string `json:"entity_id,omitempty"`
}

func (c *Client) SetLightState(lightParams LightParams, stateParam string) ([]Light, error) {

	var state string = "off"
	if stateParam == "on" {
		state = "on"
	}

	rb, err := json.Marshal(lightParams)
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
