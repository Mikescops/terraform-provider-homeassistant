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
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"media_content_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_content_type": {
				Type:     schema.TypeString,
				Optional: true,
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

	// resource id
	d.SetId(mpID)

	return diags
}

func resourceMediaPlayerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)
	media_content_id := d.Get("media_content_id").(string)
	media_content_type := d.Get("media_content_type").(string)

	o, err := c.SetMediaPlayerState(hac.MediaPlayerParams{
		ID:               id,
		MediaContentID:   media_content_id,
		MediaContentType: media_content_type,
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
