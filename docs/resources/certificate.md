---
page_title: "Credhub: Certificate"
---

# Certificate Resource

Generate a certificate based on arguments.

## Example Usage

Basic usage

```hcl
resource "credhub_certificate" "test2" {
  name               = "mycertificate"
  key_length         = 2048
  duration           = 365
  organization       = ""
  organization_unit  = ""
  locality           = ""
  state              = ""
  country            = ""
  alternative_names  = []
  key_usage          = ["digital_signature"]
  extended_key_usage = []
  ca                 = ""
  is_ca              = false
  common_name        = "example.com"
  self_sign          = true
  rotate_interval    = "30d"
  // 30 days
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of your credential
- `rotate_interval` - (Optional, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. This actually imply that a cron will recall terraform apply/plan.
- `key_length` - (Optional, Default: *Default value from credhub*) Key length of generated credential value. Values can be `2048`, `3072` or `4096`.
- **common_name**<sup>1</sup>: (Optional, Default: *NULL*) Common name of generated credential value
- `duration` - (Optional, Default: *Default value from credhub*) Duration in days of generated credential value.
- **organization**<sup>1</sup>: (Optional, Default: *NULL*) Organization of generated credential value.
- **organization_unit**<sup>1</sup>: (Optional, Default: *NULL*) Organization Unit of generated credential value.
- **locality**<sup>1</sup>: (Optional, Default: *NULL*) Locality/city of generated credential value.
- **state**<sup>1</sup>: (Optional, Default: *NULL*) State/province of generated credential value.
- **country**<sup>1</sup>: (Optional, Default: *NULL*) Country of generated credential value.
- `alternative_names` - (Optional, Default: *NULL*) Alternative names of generated credential value.
- `key_usage` - (Optional, Default: *NULL*) Key usage extensions of generated credential value. Acceptable key usages are `digital_signature`, `non_repudiation`, `key_encipherment`, `data_encipherment`, `key_agreement`, `key_cert_sign`, `crl_sign`, `encipher_only` and `decipher_only`.
- `extended_key_usage` - (Optional, Default: *NULL*) Extended key usage extensions of generated credential value. Acceptable extended key usages are `client_auth`, `server_auth`, `code_signing`, `email_protection` and `timestamping`.
- **ca**<sup>2</sup>: (Optional, Default: *NULL*) Name of certificate authority to sign of generated credential value.
- **is_ca**<sup>2</sup>: (Optional, Default: *false*) Whether to generate credential value as a certificate authority. This should be the name of a certificate credential in your credhub.
- **self_sign**<sup>2</sup>: (Optional, Default: *false*) Whether to self-sign generated credential value.

<sup>1</sup>: One subject field must be specified in the request.

<sup>2</sup>: At least one signing parameter must be provided.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
