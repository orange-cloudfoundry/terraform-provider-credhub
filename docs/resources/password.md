---
page_title: "Credhub: Password"
---

# Password Resource

Generate a password based on arguments.

## Example Usage

Basic usage

```hcl
resource "credhub_password" "mypassword" {
  name            = "mypassword"
  length          = 12
  include_special = true
  exclude_number  = false
  exclude_lower   = false
  exclude_upper   = false
  rotate_interval = "10s"
  //10 seconds
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of your credential
- `rotate_interval` - (Optional, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. This actually imply that a cron will recall terraform apply/plan.
- `length` - (Optional, Default: *Default value from credhub*) Length of generated credential value.
- `exclude_upper` - (Optional, Default: *Default value from credhub*) Exclude upper alpha characters from generated credential value.
- `exclude_lower` - (Optional, Default: *Default value from credhub*) Exclude lower alpha characters from generated credential value.
- `exclude_number` - (Optional, Default: *Default value from credhub*) Exclude number characters from generated credential value.
- `include_special` - (Optional, Default: *Default value from credhub*) Include special characters from generated credential value.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
