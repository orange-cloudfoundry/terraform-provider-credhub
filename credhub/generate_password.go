package credhub

import (
	"github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/generate"
	"github.com/hashicorp/terraform/helper/schema"
)

type GeneratePasswordResource struct {
}

func (GeneratePasswordResource) Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	creds, err := client.GeneratePassword(Name(d), generate.Password{
		Length:         d.Get("length").(int),
		IncludeSpecial: d.Get("include_special").(bool),
		ExcludeLower:   d.Get("exclude_lower").(bool),
		ExcludeNumber:  d.Get("exclude_number").(bool),
		ExcludeUpper:   d.Get("exclude_upper").(bool),
	}, credhub.Overwrite)
	if err != nil {
		return err
	}
	d.SetId(creds.Id)
	return nil
}
func (GeneratePasswordResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"exclude_lower": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"exclude_upper": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}
