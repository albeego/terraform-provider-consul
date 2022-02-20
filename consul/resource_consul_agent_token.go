package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConsulAgentToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceConsulAgentTokenCreate,
		Update: resourceConsulAgentTokenCreate,
		Read:   resourceConsulAgentTokenRead,
		Delete: resourceConsulAgentTokenDelete,

		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConsulAgentTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client, _, _ := getClient(d, meta)
	agent := client.Agent()

	token := d.Get("token").(string)
	_, err := agent.UpdateAgentACLToken(token, &consulapi.WriteOptions{
		Datacenter: "",
		Namespace:  "",
		Token:      token,
	})

	if err != nil {
		return fmt.Errorf("failed to update agent acl agent token: %v", err)
	}
	d.Set("id", token)

	return nil
}

func resourceConsulAgentTokenRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceConsulAgentTokenDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
