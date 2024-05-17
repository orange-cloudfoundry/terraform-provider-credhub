package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/values"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
)

type GenericDataSource struct {
}

func (GenericDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	credMap := make(map[string]interface{})
	err := dataSourceReadGeneric(d, meta, &credMap, func(credValue interface{}) error {
		credMap["value"] = fmt.Sprint(credValue)
		return nil
	})
	if err != nil {
		return err
	}

	for name, cred := range credMap {
		credKind := reflect.TypeOf(cred).Kind()
		if credKind == reflect.String {
			continue
		}
		if credKind == reflect.Slice || credKind == reflect.Map || credKind == reflect.Struct {
			b, _ := json.Marshal(cred)
			credMap[name] = string(b)
			continue
		}
		credMap[name] = fmt.Sprint(cred)
	}
	d.Set("credential", credMap)
	return nil
}

func (GenericDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["credential"] = &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     schema.TypeString,
	}

	return sch
}

type ValueDataSource struct {
}

func (ValueDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	data := ""
	err := dataSourceReadGeneric(d, meta, &data, func(credValue interface{}) error {
		data = fmt.Sprint(credValue)
		return nil
	})
	if err != nil {
		return err
	}
	d.Set("value", data)
	return nil
}

func (ValueDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["value"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sch
}

type JsonDataSource struct {
}

func (JsonDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	data := make(map[string]interface{})
	err := dataSourceReadGeneric(d, meta, &data, func(credValue interface{}) error {
		data["value"] = fmt.Sprint(credValue)
		return nil
	})
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	d.Set("json", string(b))
	return nil
}

func (JsonDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["json"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sch
}

type PasswordDataSource struct {
}

func (PasswordDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	password := ""
	err := dataSourceReadGeneric(d, meta, &password, func(credValue interface{}) error {
		password = fmt.Sprint(credValue)
		return nil
	})
	if err != nil {
		return err
	}
	d.Set("password", password)
	return nil
}

func (PasswordDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["password"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sch
}

type CertificateDataSource struct {
}

func (CertificateDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	data := values.Certificate{}
	err := dataSourceReadGeneric(d, meta, &data, func(credValue interface{}) error {
		return fmt.Errorf("this is not a certificate credential")
	})
	if err != nil {
		return err
	}
	d.Set("ca", data.Ca)
	d.Set("ca_name", data.CaName)
	d.Set("certificate", data.Certificate)
	d.Set("private_key", data.PrivateKey)
	return nil
}

func (CertificateDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["ca"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sch["ca_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sch["certificate"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sch["private_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	return sch
}

type RSADataSource struct {
}

func (RSADataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	data := values.RSA{}
	err := dataSourceReadGeneric(d, meta, &data, func(credValue interface{}) error {
		return fmt.Errorf("this is not a RSA credential")
	})
	if err != nil {
		return err
	}
	d.Set("public_key", data.PublicKey)
	d.Set("private_key", data.PrivateKey)
	return nil
}

func (RSADataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["public_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sch["private_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	return sch
}

type SSHDataSource struct {
}

func (SSHDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	data := values.SSH{}
	err := dataSourceReadGeneric(d, meta, &data, func(credValue interface{}) error {
		return fmt.Errorf("this is not a SSH credential")
	})
	if err != nil {
		return err
	}
	d.Set("public_key", data.PublicKey)
	d.Set("private_key", data.PrivateKey)
	return nil
}

func (SSHDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["public_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sch["private_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	return sch
}

type UserDataSource struct {
}

func (UserDataSource) DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	data := values.User{}
	err := dataSourceReadGeneric(d, meta, &data, func(credValue interface{}) error {
		return fmt.Errorf("this is not an User credential")
	})
	if err != nil {
		return err
	}
	d.Set("username", data.Username)
	d.Set("password", data.Password)
	return nil
}

func (UserDataSource) DataSourceSchema() map[string]*schema.Schema {
	sch := dataSourceSchemaGeneric()
	sch["username"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sch["password"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	return sch
}
