package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
	"strings"
	"time"
)

type Resource interface {
	Create(*schema.ResourceData, interface{}) error
	Read(*schema.ResourceData, interface{}) error
	Update(*schema.ResourceData, interface{}) error
	Schema() map[string]*schema.Schema
}

type GenerateResource interface {
	Create(*schema.ResourceData, interface{}) error
	Schema() map[string]*schema.Schema
}

type DataSource interface {
	DataSourceSchema() map[string]*schema.Schema
	DataSourceRead(*schema.ResourceData, interface{}) error
}

func LoadGenerateResource(resource GenerateResource) *schema.Resource {
	resSchema := resource.Schema()
	resSchema["name"] = &schema.Schema{
		Required: true,
		Type:     schema.TypeString,
		ForceNew: true,
	}
	resSchema["signature"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	resSchema["last_generation"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	resSchema["rotate_interval"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	resSchema["changed"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
	}
	return &schema.Resource{
		Create: CreateCreateFunc(resource.Create),
		Read:   GenerateResourceRead,
		Update: CreateCreateFunc(resource.Create),
		Delete: Delete,
		Exists: Exists,
		Schema: resSchema,
	}
}

func LoadResource(resource Resource) *schema.Resource {
	resSchema := resource.Schema()
	resSchema["name"] = &schema.Schema{
		Required: true,
		Type:     schema.TypeString,
		ForceNew: true,
	}
	return &schema.Resource{
		Create: resource.Create,
		Read:   resource.Read,
		Update: resource.Update,
		Delete: Delete,
		Exists: Exists,
		Schema: resSchema,
	}
}

func LoadDataSource(DataSource DataSource) *schema.Resource {
	return &schema.Resource{
		Read:   DataSource.DataSourceRead,
		Schema: DataSource.DataSourceSchema(),
	}
}

func Name(d *schema.ResourceData) string {

	return d.Get("name").(string)
}

func SetName(d *schema.ResourceData, value string) {

	d.Set("name", value)
}

func transformCredhubError(err error) error {
	if errResp, ok := err.(*credhub.Error); ok {
		return fmt.Errorf("%s: %s", errResp.Name, errResp.Description)
	}
	return err
}

func CreateCreateFunc(create func(d *schema.ResourceData, meta interface{}) error) func(d *schema.ResourceData, meta interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		err := create(d, meta)
		if err != nil {
			return transformCredhubError(err)
		}
		client := meta.(*credhub.CredHub)
		cred, err := client.GetById(d.Id())
		if err != nil {
			return transformCredhubError(err)
		}
		d.Set("signature", generateSignature(cred.Value))
		d.Set("last_generation", fmt.Sprintf("%d", time.Now().Unix()))
		return nil
	}
}

func Delete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	cred, err := client.GetById(d.Id())
	if err != nil {
		return err
	}
	return client.Delete(cred.Name)
}

func GenerateResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	cred, err := client.GetById(d.Id())
	if err != nil {
		return transformCredhubError(err)
	}
	if !strings.HasPrefix(d.Get("name").(string), "/") {
		cred.Name = strings.TrimPrefix(cred.Name, "/")
	}
	SetName(d, cred.Name)
	markerChanged := !d.Get("changed").(bool)
	signature := generateSignature(cred.Value)
	if d.Get("signature").(string) != signature {
		d.Set("signature", signature)
		d.Set("changed", markerChanged)
	}
	if d.Get("rotate_interval").(string) == "" {
		return nil
	}
	i, err := strconv.ParseInt(d.Get("last_generation").(string), 10, 64)
	if err != nil {
		return err
	}
	tm := time.Unix(i, 0)
	expireDuration, err := ParseDuration(d.Get("rotate_interval").(string))
	if err != nil {
		return err
	}
	if tm.Add(time.Duration(expireDuration)).Before(time.Now()) {
		d.Set("last_generation", fmt.Sprintf("%d", time.Now().Unix()))
		d.Set("changed", markerChanged)
	}
	return nil
}

func generateSignature(value interface{}) string {
	h := sha512.New()
	b, err := json.Marshal(value)
	if err != nil {
		b = []byte(fmt.Sprint(value))
	}
	h.Write(b)
	sumB := h.Sum(nil)
	return fmt.Sprintf("%x", sumB)
}

func Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*credhub.CredHub)
	var cred credentials.Credential
	var err error
	if d.Id() != "" {
		cred, err = client.GetById(d.Id())
	} else {
		cred, err = client.GetLatestVersion(Name(d))
	}
	if err != nil && strings.Contains(err.Error(), "does not exist") {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	d.SetId(cred.Id)
	return true, nil
}

func SchemaSetToStringList(set *schema.Set) []string {
	data := set.List()
	finalList := make([]string, len(data))
	for i, v := range data {
		finalList[i] = v.(string)
	}
	return finalList
}

func SchemaSetToIntList(set *schema.Set) []int {
	data := set.List()
	finalList := make([]int, len(data))
	for i, v := range data {
		finalList[i] = v.(int)
	}
	return finalList
}

func validateMapToString(mapValid map[string]bool) string {
	asList := make([]string, len(mapValid))
	i := 0
	for name, _ := range mapValid {
		asList[i] = name
		i++
	}
	return strings.Join(asList, ", ")
}
