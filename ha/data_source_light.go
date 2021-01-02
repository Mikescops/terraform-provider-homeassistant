package ha

import (
	"context"

	hac "terraform-provider-ha/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLight() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLightRead,
		Schema: map[string]*schema.Schema{
			"entity_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the resource in Home Assistant",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the light",
			},
			"attributes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Light attributes",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"brightness": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Brightness of the light from 0 to 255",
						},
						"rgb_color": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "Array of colors : red, green, blue",
						},
					},
				},
			},
		},
	}
}

func dataSourceLightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	lightID := d.Get("entity_id").(string)

	light, err := c.GetLightState(lightID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", light.State); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("attributes", flattenAttributes(light.Attributes)); err != nil {
		return diag.FromErr(err)
	}

	// resource id
	d.SetId(lightID)

	return diags
}

func flattenAttributes(attributes hac.LightAttributes) []interface{} {
	c := make(map[string]interface{})
	c["brightness"] = attributes.Brightness
	c["rgb_color"] = attributes.RgbColor

	return []interface{}{c}
}
