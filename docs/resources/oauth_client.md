---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tailscale_oauth_client Resource - terraform-provider-tailscale"
subcategory: ""
description: |-
  The oauth_client resource allows you to create OAuth clients to programmatically interact with the Tailscale API.
---

# tailscale_oauth_client (Resource)

The oauth_client resource allows you to create OAuth clients to programmatically interact with the Tailscale API.

## Example Usage

```terraform
resource "tailscale_oauth_client" "sample_client" {
  description = "sample client"
  scopes      = ["all:read"]
  tags        = ["tag:test"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scopes` (Set of String) Scopes to grant to the client. See https://tailscale.com/kb/1215/ for a list of available scopes.

### Optional

- `description` (String) A description of the key consisting of alphanumeric characters. Defaults to `""`.
- `tags` (Set of String) A list of tags that access tokens generated for the OAuth client will be able to assign to devices. Mandatory if the scopes include "devices:core" or "auth_keys".

### Read-Only

- `created_at` (String) The creation timestamp of the key in RFC3339 format
- `id` (String) The client ID, also known as the key id. Used with the client secret to generate access tokens.
- `key` (String, Sensitive) The client secret, also known as the key. Used with the client ID to generate access tokens.
- `user_id` (String) ID of the user who created this key, empty for OAuth clients created by other OAuth clients.

## Import

Import is supported using the following syntax:

```shell
# Note: Sensitive fields such as the secret key are not returned by the API and will be unset in the Terraform state after import.
terraform import tailscale_oauth_client.example k1234511CNTRL
```
