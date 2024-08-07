---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tailscale_webhook Resource - terraform-provider-tailscale"
subcategory: ""
description: |-
  The webhook resource allows you to configure webhook endpoints for your Tailscale network. See https://tailscale.com/kb/1213/webhooks for more information.
---

# tailscale_webhook (Resource)

The webhook resource allows you to configure webhook endpoints for your Tailscale network. See https://tailscale.com/kb/1213/webhooks for more information.

## Example Usage

```terraform
resource "tailscale_webhook" "sample_webhook" {
  endpoint_url  = "https://example.com/webhook/endpoint"
  provider_type = "slack"
  subscriptions = ["nodeCreated", "userDeleted"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `endpoint_url` (String) The endpoint to send webhook events to.
- `subscriptions` (Set of String) The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.

### Optional

- `provider_type` (String) The provider type of the endpoint URL. Also referred to as the 'destination' for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.

### Read-Only

- `id` (String) The ID of this resource.
- `secret` (String, Sensitive) The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.

## Import

Import is supported using the following syntax:

```shell
# Webhooks can be imported using the endpoint id, e.g.,
terraform import tailscale_webhook.sample_webhook 123456789
```
