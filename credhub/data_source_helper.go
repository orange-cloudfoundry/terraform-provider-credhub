package credhub

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceReadGeneric(d *schema.ResourceData, meta, v interface{}, onMarshalErr func(interface{})) error {
	client := meta.(*credhub.CredHub)
	if d.Get("cred_id").(string) == "" && d.Get("name").(string) == "" {
		return fmt.Errorf("Id or name must be set")
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
	d.Set("type", cred.Type)
	d.Set("name", cred.Name)

	b, err := json.Marshal(cred.Value)
	if err != nil {
		b = []byte(fmt.Sprintf(`{"value": "%s"}`, fmt.Sprint(cred.Value)))
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		onMarshalErr(cred.Value)
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
