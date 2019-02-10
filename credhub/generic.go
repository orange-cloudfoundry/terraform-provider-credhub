package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

type CredData struct {
	Value      string
	Json       map[string]interface{}
	Credential map[string]interface{}
}

func (d CredData) Check() error {
	if d.Value == "" && len(d.Json) == 0 && len(d.Credential) == 0 {
		return fmt.Errorf("data_value, data_json or data_credential has to been set.")
	}
	howMany := 0
	if d.Value != "" {
		howMany++
	}
	if len(d.Json) > 0 {
		howMany++
	}
	if len(d.Credential) > 0 {
		howMany++
	}
	if howMany > 1 {
		return fmt.Errorf("Only one of data_value, data_json or data_credential can be set.")
	}
	return nil
}
func (d CredData) CredValue() interface{} {
	if d.Value != "" {
		return d.Value
	}
	if len(d.Json) > 0 {
		return d.Json
	}
	if len(d.Credential) > 0 {
		return d.Credential
	}
	return nil
}
func ParseResourceData(d *schema.ResourceData) (CredData, error) {
	dataJson := make(map[string]interface{})
	if d.Get("data_json").(string) != "" {
		err := json.Unmarshal([]byte(d.Get("data_json").(string)), &dataJson)
		if err != nil {
			return CredData{}, err
		}
	}
	return CredData{
		Value:      d.Get("data_value").(string),
		Json:       dataJson,
		Credential: d.Get("data_credential").(map[string]interface{}),
	}, nil
}

type GenericResource struct {
}

func (r GenericResource) Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	credData, err := ParseResourceData(d)
	if err != nil {
		return err
	}
	err = credData.Check()
	if err != nil {
		return err
	}

	credType := strings.ToLower(d.Get("type").(string))
	cred, err := client.SetCredential(Name(d), credType, credData.CredValue())
	d.SetId(cred.Id)
	return nil
}

func (GenericResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"data_value": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"data_json": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"data_credential": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     schema.TypeString,
		},
	}
}

func validateKeyLength(elem interface{}, index string) ([]string, []error) {
	length := elem.(int)
	if length != 2048 && length != 3072 && length != 4096 {
		return make([]string, 0), []error{fmt.Errorf("The provided key length is not supported. Valid values include '2048', '3072' and '4096'.")}
	}
	return make([]string, 0), []error{}
}
