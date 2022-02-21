package consul

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConsulAgentToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceConsulAgentTokenCreate,
		Update: resourceConsulAgentTokenCreate,
		Read:   resourceConsulAgentTokenRead,
		Delete: resourceConsulAgentTokenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConsulAgentTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client, _, writeOptions := getClient(d, meta)
	agent := client.Agent()

	token := d.Get("token").(string)
	_, err := agent.UpdateAgentACLToken(token, writeOptions)

	if err != nil {
		return fmt.Errorf("failed to update agent acl agent token: %v", err)
	}
	return nil
}

func resourceConsulAgentTokenRead(d *schema.ResourceData, meta interface{}) error {
	token := d.Get("token").(string)
	err := d.Set("token", token)
	if err != nil {
		return fmt.Errorf("failed to read agent acl agent token: %v", err)
	}
	return nil
}

func resourceConsulAgentTokenDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
