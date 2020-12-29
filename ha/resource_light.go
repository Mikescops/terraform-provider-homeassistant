package ha

import (
	"context"

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
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceLightCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceLightUpdate(ctx, d, m)
}

func resourceLightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

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
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)
	state := d.Get("state").(string)

	o, err := c.SetLightState(LightParams{
		ID: id,
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
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("entity_id").(string)

	o, err := c.SetLightState(LightParams{
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
