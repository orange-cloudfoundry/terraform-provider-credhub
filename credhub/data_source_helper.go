package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceReadGeneric(d *schema.ResourceData, meta, v interface{}, onMarshalErr func(interface{}) error) error {
	client := meta.(*credhub.CredHub)
	if d.Get("cred_id").(string) == "" && d.Get("name").(string) == "" {
		return fmt.Errorf("ID or name must be set")
	}
	var cred credentials.Credential
	var err error
	if d.Get("cred_id").(string) != "" {
		cred, err = client.GetById(d.Get("cred_id").(string))
	} else {
		cred, err = client.GetLatestVersion(d.Get("name").(string))
	}
	if err != nil {
		return err
	}
	d.SetId(cred.Id)
	if err = d.Set("type", cred.Type); err != nil {
		return err
	}
	if err = d.Set("name", cred.Name); err != nil {
		return err
	}

	b, err := json.Marshal(cred.Value)
	if err != nil {
		b = []byte(fmt.Sprintf(`{"value": "%s"}`, fmt.Sprint(cred.Value)))
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		return onMarshalErr(cred.Value)
	}
	return nil
}
func dataSourceSchemaGeneric() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cred_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
