---
layout: default
title: ansible_group_var
parent: Resources
nav_order: 2
---
## Example Usage

```terraform
resource "ansible_group_var" "extra" {
  inventory_group_name = "db"
  key                  = "ansible_user"
  value                = "postgres"
}
```

## Argument Reference

The following arguments are suported:

- `inventory_group_name` - (Required) The group name that will be used by ansible in the dynamic inventory.
- `key` - (Required) The key name for the group variable.
- `value` - (Required) Value of the group variable.

## Attributes Reference

The following attributes are exported:

- `id` - The TFState ID field of the resource. This is built from the group name and key.
