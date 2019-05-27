---
layout: default
title: ansible_host
parent: Resources
nav_order: 3
---
## Example Usage

```terraform
resource "ansible_host" "example" {
    inventory_hostname = "example.com"
    groups = ["web"]
    vars = {
        ansible_user = "admin"
    }
}
```

## Argument Reference

The following arguments are suported:

- `inventory_hostname` - (Required) The hostname that will be used by Ansible in the dynamic inventory.
- `groups` - (Optional) A list of Ansible groups that the host will belong to.
- `vars` - (Optional) A map of string values that will be attached to the Ansible host as host variables.

## Attributes Reference

The following attributes are exported:

- `id` - The TFState ID field of the resource. This will be the same as `inventory_hostname`.
