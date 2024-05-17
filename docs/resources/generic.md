---
page_title: "Credhub: Generic"
---

# Generic Resource

Generate a generic credentials based on arguments.

This generic resource has been made to prevent any future update on credhub which is not yet implemented in this provider.

We don't recommend to use it if you can't secure your configuration, you will need to set credentials directly to your configuration.

Due to limitation to put arbitrary data inside a terraform parameter, there is 3 formats to create credentials in this resource:

To see what you can do on credhub you can have look at: http://credhub-api.cfapps.io

## Example Usage

**Data credential format**

```hcl
resource "credhub_generic" "myuser" {
  type            = "user"
  name            = "/test/myuser"
  data_credential = {
    "username" = "FQnwWoxgSrDuqDLmeLpU"
    "password" = "6mRPZB3bAfb8lRpacnXsHfDhlPqFcjH2h9YDvLpL"
  }
}
```

**Data value format**

```hcl
resource "credhub_generic" "myvalue" {
  type       = "value"
  name       = "/test/myvalue"
  data_value = "myvalue"
}
```

**Data json format**

```hcl
resource "credhub_generic" "myjson" {
  type      = "json"
  name      = "/test/myjson"
  data_json = "{\"key\": {\"elem1\": \"value\", \"elem2\": \"value2\"}}"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of your credential
- `type` - (Required) Type of your credential (see: http://credhub-api.cfapps.io/#set-credentials )
- `data_value` - (Optional, Default: *NULL*) A simple value as credential parameter. This can't be used with `data_json` or `data_credential`
- `data_credential` - (Optional, Default: *NULL*) A map with string values as credential parameter. This can't be used with `data_json` or `data_value`.
- `data_json` - (Optional, Default: *NULL*) A json string as credential parameter. This can't be used with `data_credential` or `data_value`.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
