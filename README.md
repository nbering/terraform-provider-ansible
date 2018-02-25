# Ansible Terraform Provider
A Terraform provider serving as an interop layer for an Ansible [dynamic
inventory script][1].

Read the [introductory blog post][3] for an explanation of the design
motivations behind this provider.

## Installation

To install a third-party plugin on your system, follow the [plugin installation
instructions][2] provided by Terraform.

## Terraform Configuration Example

```
resource "ansible_host" "example" {
    inventory_hostname = "example.com"
    groups = ["web"]
    vars {
        ansible_user = "admin"
    }
}

resource "ansible_group" "web" {
  inventory_group_name = "web"
  children = ["foo", "bar", "baz"]
  vars {
    foo = "bar"
    bar = 2
  }
}
```

## Alternatives and Similar Projects
### [jonmorehouse/terraform-provisioner-ansible](https://github.com/jonmorehouse/terraform-provisioner-ansible)
A Terraform Provisioner that runs Ansible-Local on a target machine at creation-time.

### [adammck/terraform-inventory](https://github.com/adammck/terraform-inventory)
A very similar solution to this one, without the Logical provider. Depends on
specific Terraform resource types, and relies heavily on cloud-providers' tag
implementations.

### [Ansible Module: Terraform](http://docs.ansible.com/ansible/devel/modules/terraform_module.html)
An Ansible module that runs Terraform plans and deployments.

## License

Contributions specific to this project are made available under the
[Mozilla Public License](./LICENSE).

Code under the `vendor/` directory is copyright of the various package owners,
and made available under their own license considerations.

[1]: https://github.com/nbering/terraform-inventory/
[2]: https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin
[3]: http://nicholasbering.ca/tools/2018/01/08/introducing-terraform-provider-ansible/
