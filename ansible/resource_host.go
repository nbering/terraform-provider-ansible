package ansible

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"inventory_hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},

			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},

			"vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceHostCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("inventory_hostname").(string))

	return nil
}
