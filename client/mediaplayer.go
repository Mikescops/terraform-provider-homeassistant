package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// MediaPlayer struct
type MediaPlayer struct {
	ID         string                `json:"entity_id,omitempty"`
	State      string                `json:"state,omitempty"`
	Attributes MediaPlayerAttributes `json:"attributes,omitempty"`
}

// MediaPlayerAttributes struct
type MediaPlayerAttributes struct {
	VolumeLevel float64 `json:"volume_level,omitempty"`
	MediaTitle  string  `json:"media_title,omitempty"`
}

// GetMediaPlayerState get state of a media_player device
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

// SetMediaPlayerParams params of SetMediaPlayerState
type SetMediaPlayerParams struct {
	ID               string `json:"entity_id,omitempty"`
	MediaContentID   string `json:"media_content_id,omitempty"`
	MediaContentType string `json:"media_content_type,omitempty"`
}

// SetMediaPlayerState set state of a media_player device
func (c *Client) SetMediaPlayerState(mpParams SetMediaPlayerParams) ([]MediaPlayer, error) {

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

// SetMediaPlayerVolumeParams params of SetMediaPlayerVolume
type SetMediaPlayerVolumeParams struct {
	ID          string  `json:"entity_id,omitempty"`
	VolumeLevel float64 `json:"volume_level,omitempty"`
}

// SetMediaPlayerVolume set volume of a media_player device
func (c *Client) SetMediaPlayerVolume(params SetMediaPlayerVolumeParams) ([]MediaPlayer, error) {

	rb, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/media_player/volume_set", c.HostURL), strings.NewReader(string(rb)))
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

// StopMediaPlayerParams params of StopMediaPlayer
type StopMediaPlayerParams struct {
	ID string `json:"entity_id,omitempty"`
}

// StopMediaPlayer stop a media_player device
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
