package resources

import (
	"github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/generate"
	"github.com/hashicorp/terraform/helper/schema"
)

type GenerateSSHResource struct {
}

func (GenerateSSHResource) Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	creds, err := client.GenerateSSH(Name(d), generate.SSH{
		Comment:   d.Get("ssh_comment").(string),
		KeyLength: d.Get("key_length").(int),
	}, true)
	if err != nil {
		return err
	}
	d.SetId(creds.Id)
	return nil
}
func (GenerateSSHResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ssh_comment": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"key_length": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validateKeyLength,
		},
	}
}
