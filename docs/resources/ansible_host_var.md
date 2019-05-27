---
layout: default
title: ansible_host_var
parent: Resources
nav_order: 4
---
## Example Usage

```terraform
resource "ansible_host_var" "extra" {
  inventory_hostname = "www.example.com"
  key                = "db_host"
  value              = "db.example.com"
}
```

## Argument Reference

The following arguments are suported:

- `inventory_hostname` - (Required) The hostname that will be used by Ansible in the dynamic inventory.
- `key` - (Required) The key name for the Ansible host variable.
- `value` - (Required) Value of the Ansible group variable.
- `variable_priority` - (Optional) Accepts an integer. Determines the order in which same-key variables will be merged by the Ansible Dynamic Inventory script to produce the final set of variables for Ansible. Higher numbers will over-write lower numbers. Default: 60

## Attributes Reference

The following attributes are exported:

- `id` - The TFState ID field of the resource. This is built from the inventory hostname and key.
