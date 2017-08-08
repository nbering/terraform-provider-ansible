package ansible

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAnsibleResource(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testAccProviders,
		Steps: []r.TestStep{
			r.TestStep{
				Config: `
			resource "ansible_host" "test" {
				inventory_hostname = "example.medstack.net"
			}

			output "host_id" {
				value = "${ansible_host.test.id}"
			}
		`,
				Check: func(s *terraform.State) error {
					rID := s.RootModule().Outputs["host_id"].Value
					if "example.medstack.net" != rID {
						return fmt.Errorf("Unexpected value for resource ID: %s", rID)
					}

					return nil
				},
			},
		},
	})
}
