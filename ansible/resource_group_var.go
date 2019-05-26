package ansible

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGroupVar() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnsibleGroupVarCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"inventory_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"variable_priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},

			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAnsibleGroupVarCreate(d *schema.ResourceData, _ interface{}) error {
	d.SetId(fmt.Sprintf("%s/%s", d.Get("inventory_group_name"), d.Get("key")))
	return nil
}
