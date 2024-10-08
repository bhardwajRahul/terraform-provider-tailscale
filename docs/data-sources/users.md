---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tailscale_users Data Source - terraform-provider-tailscale"
subcategory: ""
description: |-
  The users data source describes a list of users in a tailnet
---

# tailscale_users (Data Source)

The users data source describes a list of users in a tailnet

## Example Usage

```terraform
data "tailscale_users" "all-users" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `role` (String) Filters the users list to elements whose role is the provided value.
- `type` (String) Filters the users list to elements whose type is the provided value.

### Read-Only

- `id` (String) The ID of this resource.
- `users` (List of Object) The list of users in the tailnet (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `created` (String)
- `currently_connected` (Boolean)
- `device_count` (Number)
- `display_name` (String)
- `id` (String)
- `last_seen` (String)
- `login_name` (String)
- `profile_pic_url` (String)
- `role` (String)
- `status` (String)
- `tailnet_id` (String)
- `type` (String)
