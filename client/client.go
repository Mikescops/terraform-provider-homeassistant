package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is a strucutre containing all information about connection to host
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s, reqUrl: %s, reqBody: %s", res.StatusCode, body, req.URL, req.Body)
	}

	return body, err
}
