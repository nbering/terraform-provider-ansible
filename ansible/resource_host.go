package ansible

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostCreate,
		Read:   schema.Noop,
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
				ForceNew: true,
			},
			"vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceHostCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("inventory_hostname").(string))

	return nil
}
