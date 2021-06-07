---
page_title: "Credhub: SSH"
---

# SSH Resource

Generate a ssh based on arguments.

## Example Usage

Basic usage

```hcl
resource "credhub_ssh" "myssh" {
  name            = "myssh"
  rotate_interval = "1y"
  // 1 year
  key_length      = 2048
  ssh_comment     = ""
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of your credential
- `rotate_interval` - (Optional, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. This actually imply that a cron will recall terraform apply/plan.
- `key_length` - (Optional, Default: *Default value from credhub*) Key length of generated credential value. Values can be `2048`, `3072` or `4096`.
- `ssh_comment` - (Optional, Default: *NULL*) SSH comment of generated credential value.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
