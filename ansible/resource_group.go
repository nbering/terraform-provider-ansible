package ansible

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnsibleInventoryGroupCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"inventory_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"children": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},

			"variable_priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},

			"vars": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceAnsibleInventoryGroupCreate(d *schema.ResourceData, _ interface{}) error {
	d.SetId(d.Get("inventory_group_name").(string))
	return nil
}
