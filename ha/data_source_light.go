package ha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLight() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLightRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	lightId := d.Get("id").(string)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/states/light."+lightId, "https://home.pixelswap.fr:8888"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// this should not be here
	bearerToken := "xxxx"
	bearer := "Bearer " + bearerToken
	req.Header.Add("Authorization", bearer)

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	/** Just a debug hack to get the output of the API**/

	// body, _ := ioutil.ReadAll(r.Body)

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Warning,
	// 	Summary:  "DEBUG WARNING LOG",
	// 	Detail:   string([]byte(body)),
	// })

	// return diags

	light := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&light)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", light["state"]); err != nil {
		return diag.FromErr(err)
	}

	// resource id
	d.SetId(lightId)

	return diags
}
