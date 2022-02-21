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
			"secret_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"datacenter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"token_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Token name as found in the agent configuration: default, agent and replication.",
			},
		},
	}
}

func resourceConsulAgentTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client, _, writeOptions := getClient(d, meta)
	agent := client.Agent()

	token := d.Get("secret_id").(string)
	tokenName := d.Get("token_name").(string)

	var err error
	switch tokenName {
	case "default":
		_, err = agent.UpdateDefaultACLToken(token, writeOptions)
		break
	case "agent":
		_, err = agent.UpdateAgentACLToken(token, writeOptions)
		break
	case "replication":
		_, err = agent.UpdateReplicationACLToken(token, writeOptions)
		break
	default:
		return fmt.Errorf("failed to update acl %s token, not a valid type. Should be default, agent or replication", tokenName)
	}

	if err != nil {
		return fmt.Errorf("failed to update agent acl agent token: %v", err)
	}
	d.SetId(token)
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
