package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/generate"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type GenerateUserResource struct {
}

func (GenerateUserResource) Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	creds, err := client.GenerateUser(Name(d), generate.User{
		Username:       d.Get("username").(string),
		Length:         d.Get("length").(int),
		IncludeSpecial: d.Get("include_special").(bool),
		ExcludeNumber:  d.Get("exclude_number").(bool),
		ExcludeUpper:   d.Get("exclude_upper").(bool),
		ExcludeLower:   d.Get("exclude_lower").(bool),
	}, credhub.Overwrite)
	if err != nil {
		return err
	}
	d.SetId(creds.Id)
	return nil
}
func (GenerateUserResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"length": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"include_special": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"exclude_number": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"exclude_upper": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"exclude_lower": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}
