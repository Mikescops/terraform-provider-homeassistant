package ha

import (
	"context"
	"net/http"
	"time"

	hac "terraform-provider-ha/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bearer_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HA_BEARER_TOKEN", nil),
			},
			"host_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HA_HOST_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ha_light": resourceLight(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ha_light": dataSourceLight(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	bearerToken := d.Get("bearer_token").(string)
	hostURL := d.Get("host_url").(string)

	c := hac.Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    hostURL,
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c.Token = "Bearer " + bearerToken

	return &c, diags
}
