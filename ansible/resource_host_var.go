package ansible

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHostVar() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnsibleHostVarCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"inventory_hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"priority": {
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

func resourceAnsibleHostVarCreate(d *schema.ResourceData, _ interface{}) error {
	d.SetId(fmt.Sprintf("%s/%s", d.Get("inventory_hostname"), d.Get("key")))
	return nil
}
