package ha

import (
	"context"

	hac "terraform-provider-ha/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLight() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLightCreate,
		ReadContext:   resourceLightRead,
		UpdateContext: resourceLightUpdate,
		DeleteContext: resourceLightDelete,
		Schema: map[string]*schema.Schema{
			"entity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"brightness": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rgb_color": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func resourceLightCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceLightUpdate(ctx, d, m)
}

func resourceLightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	// resource id
	d.SetId(lightID)

	return diags
}

func resourceLightUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)
	state := d.Get("state").(string)
	brightness := d.Get("brightness").(int)
	rgb_color := d.Get("rgb_color").([]interface{})

	o, err := c.SetLightState(hac.LightParams{
		ID:         id,
		Brightness: brightness,
		RgbColor:   rgb_color,
	}, state)
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

func resourceLightDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)

	o, err := c.SetLightState(hac.LightParams{
		ID: id,
	}, "off")
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
