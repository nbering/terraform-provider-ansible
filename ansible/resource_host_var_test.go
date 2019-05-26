package ansible

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAnsibleHostVar(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAnsibleHostVarConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ansible_host_var.hostvar_1", "id", "www.example.com/foo"),
					resource.TestCheckResourceAttr(
						"ansible_host_var.hostvar_1", "inventory_hostname", "www.example.com"),
					resource.TestCheckResourceAttr(
						"ansible_host_var.hostvar_1", "key", "foo"),
					resource.TestCheckResourceAttr(
						"ansible_host_var.hostvar_1", "value", "bar"),
					resource.TestCheckResourceAttr(
						"ansible_host_var.hostvar_1", "priority", "60"),
				),
			},
		},
	})
}

const testAnsibleHostVarConfig = `
  resource "ansible_host_var" "hostvar_1" {
	  inventory_hostname = "www.example.com"
	  key                = "foo"
	  value              = "bar"
  }
`
