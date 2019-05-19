package ansible

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAnsibleGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAnsibleGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ansible_group.group_1", "id", "group_1"),
					resource.TestCheckResourceAttr(
						"ansible_group.group_1", "children.0", "foo"),
					resource.TestCheckResourceAttr(
						"ansible_group.group_1", "children.2", "baz"),
					resource.TestCheckResourceAttr(
						"ansible_group.group_1", "vars.foo", "bar"),
					resource.TestCheckResourceAttr(
						"ansible_group.group_1", "vars.bar", "2"),
				),
			},
		},
	})
}

const testAnsibleGroupConfig = `
  resource "ansible_group" "group_1" {
    inventory_group_name = "group_1"
    children = ["foo", "bar", "baz"]
    vars = {
      foo = "bar"
      bar = 2
    }
  }
`
