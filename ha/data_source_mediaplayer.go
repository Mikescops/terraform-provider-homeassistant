package ha

import (
	"context"

	hac "terraform-provider-ha/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMediaPlayer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMediaPlayerRead,
		Schema: map[string]*schema.Schema{
			"entity_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the resource in Home Assistant",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the device",
			},
			"attributes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Device attributes",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_level": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Volume level",
						},
						"media_title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Title of the media displayed",
						},
					},
				},
			},
		},
	}
}

func dataSourceMediaPlayerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceID := d.Get("entity_id").(string)

	mediaplayer, err := c.GetMediaPlayerState(resourceID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", mediaplayer.State); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("attributes", flattenMediaPlayerAttributes(mediaplayer.Attributes)); err != nil {
		return diag.FromErr(err)
	}

	// resource id
	d.SetId(resourceID)

	return diags
}

func flattenMediaPlayerAttributes(attributes hac.MediaPlayerAttributes) []interface{} {
	c := make(map[string]interface{})
	c["volume_level"] = attributes.VolumeLevel
	c["media_title"] = attributes.MediaTitle

	return []interface{}{c}
}
