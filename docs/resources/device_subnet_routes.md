---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tailscale_device_subnet_routes Resource - terraform-provider-tailscale"
subcategory: ""
description: |-
  The device_subnet_routes resource allows you to configure enabled subnet routes for your Tailscale devices. See https://tailscale.com/kb/1019/subnets for more information.
  Routes must be both advertised and enabled for a device to act as a subnet router or exit node. Routes must be advertised directly from the device: advertised routes cannot be managed through Terraform. If a device is advertising routes, they are not exposed to traffic until they are enabled. Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question is advertising them.
  Note: all routes enabled for the device through the admin console or autoApprovers in the ACL must be explicitly added to the routes attribute of this resource to avoid configuration drift.
---

# tailscale_device_subnet_routes (Resource)

The device_subnet_routes resource allows you to configure enabled subnet routes for your Tailscale devices. See https://tailscale.com/kb/1019/subnets for more information.

Routes must be both advertised and enabled for a device to act as a subnet router or exit node. Routes must be advertised directly from the device: advertised routes cannot be managed through Terraform. If a device is advertising routes, they are not exposed to traffic until they are enabled. Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question is advertising them.

Note: all routes enabled for the device through the admin console or autoApprovers in the ACL must be explicitly added to the routes attribute of this resource to avoid configuration drift.

## Example Usage

```terraform
data "tailscale_device" "sample_device" {
  name = "device.example.com"
}

resource "tailscale_device_subnet_routes" "sample_routes" {
  # Prefer the new, stable `node_id` attribute; the legacy `.id` field still works.
  device_id = data.tailscale_device.sample_device.node_id
  routes = [
    "10.0.1.0/24",
    "1.2.0.0/16",
    "2.0.0.0/24"
  ]
}

resource "tailscale_device_subnet_routes" "sample_exit_node" {
  # Prefer the new, stable `node_id` attribute; the legacy `.id` field still works.
  device_id = data.tailscale_device.sample_device.node_id
  routes = [
    # Configure as an exit node
    "0.0.0.0/0",
    "::/0"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) The device to set subnet routes for
- `routes` (Set of String) The subnet routes that are enabled to be routed by a device

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Device subnet rules can be imported using the node ID (preferred), e.g.,
terraform import tailscale_device_subnet_routes.sample nodeidCNTRL
# Device subnet rules can be imported using the legacy ID, e.g.,
terraform import tailscale_device_subnet_routes.sample 123456789
```
