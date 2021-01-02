package ha

import (
	"context"

	hac "terraform-provider-ha/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMediaPlayer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMediaPlayerCreate,
		ReadContext:   resourceMediaPlayerRead,
		UpdateContext: resourceMediaPlayerUpdate,
		DeleteContext: resourceMediaPlayerDelete,
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
			"media_content_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the content to display",
			},
			"media_content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the content to display",
			},
			"volume_level": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "Volume level from 0.01 to 1",
			},
		},
	}
}

func resourceMediaPlayerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceMediaPlayerUpdate(ctx, d, m)
}

func resourceMediaPlayerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	mpID := d.Get("entity_id").(string)

	mediaplayer, err := c.GetMediaPlayerState(mpID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", mediaplayer.State); err != nil {
		return diag.FromErr(err)
	}

	volumeLevelTruncated := toFixed(mediaplayer.Attributes.VolumeLevel, 1)
	if err := d.Set("volume_level", volumeLevelTruncated); err != nil {
		return diag.FromErr(err)
	}

	// resource id
	d.SetId(mpID)

	return diags
}

func resourceMediaPlayerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)
	mediaContentID := d.Get("media_content_id").(string)
	mediaContentType := d.Get("media_content_type").(string)

	o, err := c.SetMediaPlayerState(hac.SetMediaPlayerParams{
		ID:               id,
		MediaContentID:   mediaContentID,
		MediaContentType: mediaContentType,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if len(o) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "No change applied",
			Detail:   "No light was updated by your Terraform changes",
		})
		d.SetId(id)
	} else {
		d.SetId(o[0].ID)
	}

	volumeLevel, isVolumeSet := d.GetOk("volume_level")
	volumeLevelTruncated := toFixed(volumeLevel.(float64), 2)
	if isVolumeSet {
		_, err := c.SetMediaPlayerVolume(hac.SetMediaPlayerVolumeParams{
			ID:          id,
			VolumeLevel: volumeLevelTruncated,
		})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceMediaPlayerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)

	o, err := c.StopMediaPlayer(hac.StopMediaPlayerParams{
		ID: id,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if len(o) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "No change applied",
			Detail:   "No light was updated by your Terraform changes",
		})
	}

	d.SetId("")

	return diags
}
