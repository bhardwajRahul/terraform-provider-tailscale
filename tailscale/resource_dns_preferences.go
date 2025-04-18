// Copyright (c) David Bond, Tailscale Inc, & Contributors
// SPDX-License-Identifier: MIT

package tailscale

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"tailscale.com/client/tailscale/v2"
)

func resourceDNSPreferences() *schema.Resource {
	return &schema.Resource{
		Description:   "The dns_preferences resource allows you to configure DNS preferences for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.",
		ReadContext:   resourceDNSPreferencesRead,
		CreateContext: resourceDNSPreferencesCreate,
		UpdateContext: resourceDNSPreferencesUpdate,
		DeleteContext: resourceDNSPreferencesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"magic_dns": {
				Type:        schema.TypeBool,
				Description: "Whether or not to enable magic DNS",
				Required:    true,
			},
		},
	}
}

func resourceDNSPreferencesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tailscale.Client)

	preferences, err := client.DNS().Preferences(ctx)
	if err != nil {
		return diagnosticsError(err, "Failed to fetch dns preferences")
	}

	if err = d.Set("magic_dns", preferences.MagicDNS); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceDNSPreferencesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tailscale.Client)
	magicDNS := d.Get("magic_dns").(bool)
	preferences := tailscale.DNSPreferences{
		MagicDNS: magicDNS,
	}

	if err := client.DNS().SetPreferences(ctx, preferences); err != nil {
		return diagnosticsError(err, "Failed to set dns preferences")
	}

	d.SetId(createUUID())
	return resourceDNSPreferencesRead(ctx, d, m)
}

func resourceDNSPreferencesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("magic_dns") {
		return resourceDNSPreferencesRead(ctx, d, m)
	}

	client := m.(*tailscale.Client)
	magicDNS := d.Get("magic_dns").(bool)

	preferences := tailscale.DNSPreferences{
		MagicDNS: magicDNS,
	}

	if err := client.DNS().SetPreferences(ctx, preferences); err != nil {
		return diagnosticsError(err, "Failed to set dns preferences")
	}

	return resourceDNSPreferencesRead(ctx, d, m)
}

func resourceDNSPreferencesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tailscale.Client)

	if err := client.DNS().SetPreferences(ctx, tailscale.DNSPreferences{}); err != nil {
		return diagnosticsError(err, "Failed to set dns preferences")
	}

	return nil
}
