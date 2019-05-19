package ansible

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"ansible_host":     resourceHost(),
			"ansible_host_var": resourceHostVar(),
			"ansible_group":    resourceGroup(),
		},
	}
}
