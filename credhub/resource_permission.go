package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/orange-cloudfoundry/terraform-provider-credhub/credhub/extend"
)

func resourcePermission() *schema.Resource {

	return &schema.Resource{
		Create: resourcePermissionCreate,
		Read:   resourcePermissionRead,
		Update: resourcePermissionUpdate,
		Delete: resourcePermissionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"actor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"operations": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice(
						[]string{"read", "write", "delete", "write_acl", "read_acl"},
						false),
				},
				Set: schema.HashString,
			},
		},
	}
}

func resourcePermissionData(d *schema.ResourceData) (path string, actor string, operations []string) {
	path = d.Get("path").(string)
	actor = d.Get("actor").(string)
	operationsI := d.Get("operations").(*schema.Set).List()

	operations = make([]string, len(operationsI))
	for i := 0; i < len(operationsI); i++ {
		operations[i] = operationsI[i].(string)
	}
	return
}

func resourcePermissionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*credhub.CredHub)

	path, actor, operations := resourcePermissionData(d)

	permission, err := client.AddPermission(path, actor, operations)
	if err != nil {
		return err
	}
	d.SetId(permission.UUID)
	return nil
}

func resourcePermissionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*credhub.CredHub)
	permExtend := extend.NewPermission(client)
	path, actor, operations := resourcePermissionData(d)

	_, err = permExtend.UpdatePermission(d.Id(), path, actor, operations)
	if err != nil {
		return err
	}
	return nil
}

func resourcePermissionRead(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*credhub.CredHub)

	permission, err := client.GetPermission(d.Id())
	if err != nil {
		return err
	}

	d.Set("path", permission.Path)
	d.Set("actor", permission.Actor)
	d.Set("operations", permission.Operations)

	return nil
}

func resourcePermissionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*credhub.CredHub)
	permExtend := extend.NewPermission(client)

	return permExtend.DeletePermission(d.Id())
}
