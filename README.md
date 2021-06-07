# terraform-provider-credhub  [![Build Status](https://travis-ci.org/orange-cloudfoundry/terraform-provider-credhub.svg?branch=master)](https://travis-ci.org/orange-cloudfoundry/terraform-provider-credhub)

This terraform provider lets you create and retrieve credentials from [credhub](https://github.com/cloudfoundry-incubator/credhub).

## Installations

**Requirements:** You need, of course, terraform (**>=0.13**) which is available here: https://www.terraform.io/downloads.html

Add to your terraform file:

```hcl
terraform {
  required_providers {
    cfsecurity = {
      source  = "orange-cloudfoundry/credhub"
      version = "latest"
    }
  }
}
```

## Documentation

You can find documentation at https://registry.terraform.io/providers/orange-cloudfoundry/credhub/latest/docs