---
page_title: "Credhub: RSA"
---

# RSA Resource

Generate a rsa based on arguments.

## Example Usage

Basic usage

```hcl
resource "credhub_rsa" "myrsa" {
  name            = "myrsa"
  rotate_interval = "4w"
  // 4 weeks
  key_length      = 2048
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of your credential
- `rotate_interval` - (Optional, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is a helper to do password rotation. This actually imply that a cron will recall terraform apply/plan.
- `key_length` - (Optional, Default: *Default value from credhub*) Key length of generated credential value. Values can be `2048`, `3072` or `4096`.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
