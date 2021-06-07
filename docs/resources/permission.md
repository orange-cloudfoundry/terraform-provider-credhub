---
page_title: "Credhub: Permission"
---

# Permission Resource

Manage permission for a or many credentials.

See authentication-specific identities [explained here](https://github.com/cloudfoundry-incubator/credhub/blob/master/docs/authentication-identities.md)

## Example Usage

Basic usage

```hcl
resource "credhub_permission" "myuser_perm_mypath" {
  path       = "/a/path/*"
  actor      = "uaa-user:dc912b22-caeb-4780-a6d5-aa5843f81868"
  operations = ["read", "write", "delete"]
}
```

## Argument Reference

The following arguments are supported:

- `path` - (Required) A path where you would like to add a permission to for an actor
- `actor` - (Required) An actor that receives permission at the specified path
  (See authentication-specific identities [explained here](https://github.com/cloudfoundry-incubator/credhub/blob/master/docs/authentication-identities.md))
- `operations` - (Required) List of operations given to actor for specified path
  (supported operations: `read`, `write`, `delete`, `read_acl`, `write_acl`)

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
