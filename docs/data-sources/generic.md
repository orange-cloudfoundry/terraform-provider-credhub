---
page_title: "Credhub: Generic"
---

# Generic Data Source

Retrieve generic type.

## Example Usage

```hcl
data "credhub_generic" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required if cred_id empty) Name of your credential.
- `cred_id`: (Required if name empty) Your credential ID.

## Attributes Reference

The following attributes are exported:

- `type` - This is the type of your credential
- `credential` - Data retrieve from your credential. This is a map of string, a type value will be accessible by `credential.value`, data with a subtree will be converted to json if needed.