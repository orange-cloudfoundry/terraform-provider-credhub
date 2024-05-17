---
page_title: "Credhub: Certificate"
---

# Certificate Data Source

Retrieve certificate type.

## Example Usage

```hcl
data "credhub_certificate" "my_data" {
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

- `type` - This is the type of your credential but here it will be always `certificate`.
- `ca` - CA from your credential.
- `ca_name` - CA Name from your credential.
- `certificate` - Certificate in pem format from your credential.
- `private_key` - Private key in pem format from your credential.