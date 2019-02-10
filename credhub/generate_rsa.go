package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/generate"
	"github.com/hashicorp/terraform/helper/schema"
)

type GenerateRSAResource struct {
}

func (GenerateRSAResource) Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	creds, err := client.GenerateRSA(Name(d), generate.RSA{
		KeyLength: d.Get("key_length").(int),
	}, credhub.Overwrite)
	if err != nil {
		return err
	}
	d.SetId(creds.Id)
	return nil
}

// '2048', '3072' and '4096'
func (GenerateRSAResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key_length": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validateKeyLength,
		},
	}
}
