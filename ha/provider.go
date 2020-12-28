package ha

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bearer_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HA_BEARER_TOKEN", nil),
			},
			"host_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HA_HOST_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"ha_light": dataSourceLight(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	bearerToken := d.Get("bearer_token").(string)
	hostURL := d.Get("host_url").(string)

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    hostURL,
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c.Token = "Bearer " + bearerToken

	return &c, diags
}

/*** Following should be in a proper client module ***/

type Light struct {
	ID    string `json:"entity_id,omitempty"`
	State string `json:"state,omitempty"`
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
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func (c *Client) GetLightState(lightID string) (*Light, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/states/light.%s", c.HostURL, lightID), nil)
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
