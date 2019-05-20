---
layout: home
---
# Ansible Terraform Provider
A Terraform provider serving as an interop layer for an Ansible [dynamic
inventory script][1].

Read the [introductory blog post][3] for an explanation of the design
motivations behind this provider.

## Terraform Configuration Example

{% highlight tf %}
resource "ansible_host" "example" {
    inventory_hostname = "example.com"
    groups = ["web"]
    vars = {
        ansible_user = "admin"
    }
}

resource "ansible_group" "web" {
  inventory_group_name = "web"
  children = ["foo", "bar", "baz"]
  vars = {
    foo = "bar"
    bar = 2
  }
}
{% endhighlight %}

## Compatibility

Version [1.0.0][7] of this project is compatible with Terraform version
[0.12-beta2][8]. You will also need [2.0.0][11]+ of the terraform-inventory
script, as the internal structure of Terraform state files has changed. 

If you need a version compatible with an earlier version of
Terraform, use version [0.0.4][9].

When upgrading to Terraform 0.12.x, you may need to change your configuration
files to reflect changes to the new version of the Hashicorp Configuration
Lanaguage (HCL). The only known incompatibility is that `vars` attributes now
require an equals operator (`=`).

**0.11.x**
{% highlight tf %}
resource "ansible_host" "example" {
    inventory_hostname = "example.com"
    vars {
        ansible_user = "admin"
    }
}
{% endhighlight %}

**0.12.x**
{% highlight tf %}
resource "ansible_host" "example" {
    inventory_hostname = "example.com"
    vars = {
        ansible_user = "admin"
    }
}
{% endhighlight %}

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
[4]: https://github.com/nbering/terraform-provider-ansible/releases
[5]: https://golang.org/doc/install
[6]: https://github.com/travis-ci/gimme
[7]: https://github.com/nbering/terraform-provider-ansible/releases/tag/v1.0.0
[8]: https://github.com/hashicorp/terraform/releases/tag/v0.12.0-beta2
[9]: https://github.com/nbering/terraform-provider-ansible/releases/tag/v0.0.4
[10]: https://bazaar.canonical.com/en/
[11]: https://github.com/nbering/terraform-inventory/releases/tag/v2.0.0
