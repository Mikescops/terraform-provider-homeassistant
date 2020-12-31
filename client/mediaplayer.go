package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type MediaPlayer struct {
	ID         string                `json:"entity_id,omitempty"`
	State      string                `json:"state,omitempty"`
	Attributes MediaPlayerAttributes `json:"attributes,omitempty"`
}

type MediaPlayerAttributes struct {
	VolumeLevel float32 `json:"volume_level,omitempty"`
	MediaTitle  string  `json:"media_title,omitempty"`
}

func (c *Client) GetMediaPlayerState(mpID string) (*MediaPlayer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/states/%s", c.HostURL, mpID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	mediaplayer := MediaPlayer{}
	err = json.Unmarshal(body, &mediaplayer)
	if err != nil {
		return nil, err
	}

	return &mediaplayer, nil
}

type MediaPlayerParams struct {
	ID               string `json:"entity_id,omitempty"`
	MediaContentID   string `json:"media_content_id,omitempty"`
	MediaContentType string `json:"media_content_type,omitempty"`
}

func (c *Client) SetMediaPlayerState(mpParams MediaPlayerParams) ([]MediaPlayer, error) {

	rb, err := json.Marshal(mpParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/media_extractor/play_media", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	mediaplayer := []MediaPlayer{}
	err = json.Unmarshal(body, &mediaplayer)
	if err != nil {
		return nil, err
	}

	return mediaplayer, nil
}

type StopMediaPlayerParams struct {
	ID string `json:"entity_id,omitempty"`
}

func (c *Client) StopMediaPlayer(params StopMediaPlayerParams) ([]MediaPlayer, error) {

	rb, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/media_player/media_stop", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	mediaplayer := []MediaPlayer{}
	err = json.Unmarshal(body, &mediaplayer)
	if err != nil {
		return nil, err
	}

	return mediaplayer, nil
}
