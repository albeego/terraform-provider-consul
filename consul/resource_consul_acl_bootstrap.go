package consul

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceConsulACLBootstrap() *schema.Resource {
	return &schema.Resource{
		Create: resourceConsulACLBootstrapCreate,
		Read:   resourceConsulACLBootstrapRead,
		Delete: resourceConsulACLBootstrapDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Computed:    true,
				Optional:    true,
				Sensitive:   true,
				Description: "The secret id.",
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConsulACLBootstrapCreate(d *schema.ResourceData, meta interface{}) error {
	client, _, _ := getClient(d, meta)

	log.Printf("[DEBUG] Creating ACL Bootstrap")

	token, _, err := client.ACL().Bootstrap()
	if err != nil {
		return fmt.Errorf("error creating ACL token: %s", err)
	}

	log.Printf("[DEBUG] Created ACL token %q", token.AccessorID)

	d.SetId(token.AccessorID)
	d.Set("secret_id", token.SecretID)
	d.Set("token", token.SecretID)

	return nil
}

func resourceConsulACLBootstrapRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceConsulACLBootstrapDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
