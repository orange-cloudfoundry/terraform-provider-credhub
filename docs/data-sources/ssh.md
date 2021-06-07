---
page_title: "Credhub: SSH"
---

# SSH Data Source

Retrieve ssh type.

## Example Usage

```hcl
data "credhub_ssh" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required if cred_id empty) Name of your credential.
- `cred_id`: (Required if name empty) Id of your credential.

## Attributes Reference

The following attributes are exported:

- `type` - This is the type of your credential but here it will be always `ssh`.
- `public_key` - Public key from your credential.
- `private_key` - Private key from your credential.