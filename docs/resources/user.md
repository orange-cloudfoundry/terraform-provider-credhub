---
page_title: "Credhub: User"
---

# User Resource

Generate a user based on arguments.

## Example Usage

Basic usage

```hcl
resource "credhub_user" "myuser" {
  name            = "myuser"
  username        = "labelette"
  length          = 12
  include_special = true
  exclude_number  = false
  exclude_lower   = false
  exclude_upper   = false
  rotate_interval = "5h"
  //5 hours
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of your credential
- `rotate_interval` - (Optional, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is a helper to do password rotation. This actually imply that a cron will recall terraform apply/plan.
- `username` - (Required) User provided value for username.
- `length` - (Optional, Default: *Default value from credhub*) Length of generated credential value.
- `exclude_upper` - (Optional, Default: *Default value from credhub*) Exclude upper alpha characters from generated credential value.
- `exclude_lower` - (Optional, Default: *Default value from credhub*) Exclude lower alpha characters from generated credential value.
- `exclude_number` - (Optional, Default: *Default value from credhub*) Exclude number characters from generated credential value.
- `include_special` - (Optional, Default: *Default value from credhub*) Include special characters from generated credential value.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
