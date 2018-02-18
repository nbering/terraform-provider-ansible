package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/nbering/terraform-provider-ansible/ansible"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ansible.Provider,
	})
}
