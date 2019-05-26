package ansible

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAnsibleGroupVar(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAnsibleGroupVarConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ansible_group_var.groupvar_1", "id", "web/baz"),
					resource.TestCheckResourceAttr(
						"ansible_group_var.groupvar_1", "inventory_group_name", "web"),
					resource.TestCheckResourceAttr(
						"ansible_group_var.groupvar_1", "key", "baz"),
					resource.TestCheckResourceAttr(
						"ansible_group_var.groupvar_1", "value", "bup"),
					resource.TestCheckResourceAttr(
						"ansible_group_var.groupvar_1", "variable_priority", "60"),
				),
			},
		},
	})
}

const testAnsibleGroupVarConfig = `
  resource "ansible_group_var" "groupvar_1" {
	  inventory_group_name = "web"
	  key                  = "baz"
	  value                = "bup"
  }
`
