---
page_title: "Provider: CredHub provider"
description: Provider for using CredHub API.
---

# CredHub Provider

Provider for using [credhub](https://github.com/cloudfoundry-incubator/credhub) api, this will let you interact with api to create and retrieve secrets from credhub.

You will be able to use this provider alongside cloud foundry provider.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "credhub" {
  credhub_server      = "https://api.of.your.credhub.com"
  username            = "user"
  password            = "mypassword"
  skip_ssl_validation = false
  client_id           = ""
  client_secret       = ""
  ca_cert             = ""
}
```

## Argument Reference

The following arguments are supported:

- `credhub_server` - (Required) Your credhub api url. You can use env var `CREDHUB_SERVER`.
- `username` - (Optional, default: `null`) The username of a UAA user `credhub.write` and `credhub.read` scopes. You can use env var `CREDHUB_USERNAME`. (Optional if you use an client_id/client_secret)
- `password` - (Optional, default: `null`) The password of a UAA user `credhub.write` and `credhub.read` scopes. You can use env var `CREDHUB_PASSWORD`. (Optional if you use an client_id/client_secret)
- `skip_ssl_validation` - (Optional, default: `null`) Set to true to skip verification of the API endpoint. Not recommended!
- `client_id` - (Optional, default: `null`) The client_id of a UAA client `credhub.write` and `credhub.read` scopes. You can use env var `CREDHUB_CLIENT`. (Optional if you use a username/password)
- `client_secret` - (Optional, default: `null`) The client_secret of a UAA client `credhub.write` and `credhub.read` scopes. You can use env var `CREDHUB_SECRET`. (Optional if you use a username/password)
- `ca_cert` - (Optional, default: `null`) Trusted CA for API and UAA TLS connections. You can use env var `CREDHUB_CA_CERT`.

