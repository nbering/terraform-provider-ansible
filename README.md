# Ansible Terraform Provider
A Terraform provider serving as an interop layer for an Ansible [dynamic
inventory script][1].

## Terraform Configuration Example

```
resource "ansible_host" "example" {
    inventory_hostname = "example.com"
    groups = ["web"]
    vars {
        ansible_user = "admin"
    }
}
```

## License

Contributions specific to this project are made available under the
[Mozilla Public License](./LICENSE).

Code under the `vendor/` directory is copyright of the various package owners,
and made available under their own license considerations.

[1]: https://gitlab.com/nbering/terraform-inventory/
