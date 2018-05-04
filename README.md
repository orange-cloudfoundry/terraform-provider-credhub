# terraform-provider-credhub  [![Build Status](https://travis-ci.org/orange-cloudfoundry/terraform-provider-credhub.svg?branch=master)](https://travis-ci.org/orange-cloudfoundry/terraform-provider-credhub)

This terraform provider lets you create and retrieve credentials from [credhub](https://github.com/cloudfoundry-incubator/credhub).


## Installations

**Requirements:** You need, of course, terraform (**>=0.8**) which is available here: https://www.terraform.io/downloads.html

### Automatic

To install a specific version, set PROVIDER_CREDHUB_VERSION before executing the following command

```bash
$ export PROVIDER_CREDHUB_VERSION="v0.10.0"
```

#### via curl

```bash
$ bash -c "$(curl -fsSL https://raw.github.com/orange-cloudfoundry/terraform-provider-credhub/master/bin/install.sh)"
```

#### via wget

```bash
$ bash -c "$(wget https://raw.github.com/orange-cloudfoundry/terraform-provider-credhub/master/bin/install.sh -O -)"
```

### Manually

1. Get the build for your system in releases: https://github.com/orange-cloudfoundry/terraform-provider-credhub/releases/latest
2. Create a `providers` directory inside terraform user folder: `mkdir -p ~/.terraform.d/providers`
3. Move the provider previously downloaded in this folder: `mv /path/to/download/directory/terraform-provider-credhub ~/.terraform.d/providers`
4. Ensure provider is executable: `chmod +x ~/.terraform.d/providers/terraform-provider-credhub`
5. add `providers` path to your `.terraformrc`:
```bash
cat <<EOF > ~/.terraformrc
providers {
    credhub = "/full/path/to/.terraform.d/providers/terraform-provider-credhub"
}
EOF
```

6. you can now perform any terraform action on [credhub](https://github.com/cloudfoundry-incubator/credhub) resources

## provider configuration

```tf
provider "credhub" {
  credhub_server = "https://api.of.your.credhub.com"
  username = "user"
  password = "mypassword"
  skip_ssl_validation = false
  client_id = ""
  client_secret = ""
  ca_cert = ""
}
```

- **credhub_server**: (**Required**, *Env Var: `CREDHUB_SERVER`*) Your credhub api url.
- **username**: *(Optional, default: `null`, Env Var: `CREDHUB_USERNAME`)* The username of an UAA user credhub.write and credhub.read scopes. (Optional if you use an client_id/client_secret)
- **password**: *(Optional, default: `null`, Env Var: `CREDHUB_PASSWORD`)* The password of an UAA user credhub.write and credhub.read scopes. (Optional if you use an client_id/client_secret)
- **skip_ssl_validation**: *(Optional, default: `false`)* Set to true to skip verification of the API endpoint. Not recommended!.
- **client_id**: *(Optional, default: `null`, Env Var: `CREDHUB_CLIENT`)* The client_id of an UAA client credhub.write and credhub.read scopes. (Optional if you use an username/password)
- **client_secret**: *(Optional, default: `null`, Env Var: `CREDHUB_SECRET`)* The client_secret of an UAA client credhub.write and credhub.read scopes. (Optional if you use an username/password)
- **ca_cert**: *(Optional, default: `null`, Env Var: `CREDHUB_CA_CERT`))* Trusted CA for API and UAA TLS connections.

## Resources and Data sources

- [Resources](#resources)
  - [Generate password](#generate-password)
  - [Generate certificate](#generate-certificate)
  - [Generate RSA](#generate-rsa)
  - [Generate SSH](#generate-ssh)
  - [Generate user](#generate-user)
  - [Generic](#generic)
- [Datasources](#datasources)
  - [Value](#value)
  - [Json](#json)
  - [Password](#password)
  - [Certificate](#certificate)
  - [RSA](#rsa)
  - [SSH](#ssh)
  - [User](#user)
  - [Generic](#generic-1)

## Resources

**IMPORTANT**: Credhub-generated resources are most secure as you will never have to set confidential data as plain text in TF specs (in comparison to use of the `credhub_generic` resource). Data from Credhub-generated resources are never stored in your tfstate either: this provider stores instead a fingerprint of this data in the tfstate, and uses this fingerprint to determine if it should be updated or not.

### Generate password

```hcl
resource "credhub_password" "mypassword" {
  name = "mypassword"
  length = 12
  include_special = true
  exclude_number = false
  exclude_lower = false
  exclude_upper = false
  rotate_interval = "10s" //10 seconds
}
```

- **name**: (**Required**) Name of your credential
- **rotate_interval**: (*Optional*, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. 
This actually imply that a cron will recall terraform apply/plan.
- **length**: (*Optional*, Default: *Default value from credhub*) Length of generated credential value.
- **exclude_upper**: (*Optional*, Default: *Default value from credhub*) Exclude upper alpha characters from generated credential value.
- **exclude_lower**: (*Optional*, Default: *Default value from credhub*) Exclude lower alpha characters from generated credential value.
- **exclude_number**: (*Optional*, Default: *Default value from credhub*) Exclude number characters from generated credential value.
- **include_special**: (*Optional*, Default: *Default value from credhub*) Include special characters from generated credential value.

---

### Generate certificate

```hcl
resource "credhub_certificate" "test2" {
  name = "mycertificate"
  key_length = 2048
  duration = 365
  organization = ""
  organization_unit = ""
  locality = ""
  state = ""
  country = ""
  alternative_names = []
  key_usage = ["digital_signature"]
  extended_key_usage = []
  ca = ""
  is_ca = false
  common_name = "example.com"
  self_sign = true
  rotate_interval = "30d" // 30 days
}
```

- **name**: (**Required**) Name of your credential
- **rotate_interval**: (*Optional*, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. 
This actually imply that a cron will recall terraform apply/plan.
- **key_length**: (*Optional*, Default: *Default value from credhub*) Key length of generated credential value. Values can be `2048`, `3072` or `4096`.
- **common_name**<sup>1</sup>: (*Optional*, Default: *NULL*) Common name of generated credential value
- **duration**: (*Optional*, Default: *Default value from credhub*) Duration in days of generated credential value.
- **organization**<sup>1</sup>: (*Optional*, Default: *NULL*) Organization of generated credential value.
- **organization_unit**<sup>1</sup>: (*Optional*, Default: *NULL*) Organization Unit of generated credential value.
- **locality**<sup>1</sup>: (*Optional*, Default: *NULL*) Locality/city of generated credential value.
- **state**<sup>1</sup>: (*Optional*, Default: *NULL*) State/province of generated credential value.
- **country**<sup>1</sup>: (*Optional*, Default: *NULL*) Country of generated credential value.
- **alternative_names**: (*Optional*, Default: *NULL*) Alternative names of generated credential value.
- **key_usage**: (*Optional*, Default: *NULL*) Key usage extensions of generated credential value. 
Acceptable key usages are `digital_signature`, `non_repudiation`, `key_encipherment`, `data_encipherment`, `key_agreement`, `key_cert_sign`, `crl_sign`, `encipher_only` and `decipher_only`.
- **extended_key_usage**: (*Optional*, Default: *NULL*) Extended key usage extensions of generated credential value. 
 Acceptable extended key usages are `client_auth`, `server_auth`, `code_signing`, `email_protection` and `timestamping`.
- **ca**<sup>2</sup>: (*Optional*, Default: *NULL*) Name of certificate authority to sign of generated credential value.
- **is_ca**<sup>2</sup>: (*Optional*, Default: *false*) Whether to generate credential value as a certificate authority. This should be the name of a certificate credential in your credhub.
- **self_sign**<sup>2</sup>: (*Optional*, Default: *false*) Whether to self-sign generated credential value.

<sup>1</sup>: One subject field must be specified in the request.

<sup>2</sup>: At least one signing parameter must be provided.

---

### Generate RSA

```hcl
resource "credhub_rsa" "myrsa" {
  name = "myrsa"
  rotate_interval = "4w" // 4 weeks
  key_length = 2048
}
```

- **name**: (**Required**) Name of your credential
- **rotate_interval**: (*Optional*, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. 
This actually imply that a cron will recall terraform apply/plan.
- **key_length**: (*Optional*, Default: *Default value from credhub*) Key length of generated credential value. Values can be `2048`, `3072` or `4096`.

---

### Generate SSH

```hcl
resource "credhub_ssh" "myssh" {
  name = "myssh"
  rotate_interval = "1y" // 1 year
  key_length = 2048
  ssh_comment = ""
}
```

- **name**: (**Required**) Name of your credential
- **rotate_interval**: (*Optional*, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. 
This actually imply that a cron will recall terraform apply/plan.
- **key_length**: (*Optional*, Default: *Default value from credhub*) Key length of generated credential value. Values can be `2048`, `3072` or `4096`.
- **ssh_comment**: (*Optional*, Default: *NULL*) SSH comment of generated credential value.

---

### Generate user

```hcl
resource "credhub_user" "myuser" {
  name = "myuser"
  username = "labelette"
  length = 12
  include_special = true
  exclude_number = false
  exclude_lower = false
  exclude_upper = false
  rotate_interval = "5h" //5 hours
}
```

- **name**: (**Required**) Name of your credential
- **rotate_interval**: (*Optional*, Default: *Null*) If defined provider will mark the resource as a dirty to regenerate a new password, this is an helper to do password rotation. 
This actually imply that a cron will recall terraform apply/plan.
- **username**: (**Required**) User provided value for username.
- **length**: (*Optional*, Default: *Default value from credhub*) Length of generated credential value.
- **exclude_upper**: (*Optional*, Default: *Default value from credhub*) Exclude upper alpha characters from generated credential value.
- **exclude_lower**: (*Optional*, Default: *Default value from credhub*) Exclude lower alpha characters from generated credential value.
- **exclude_number**: (*Optional*, Default: *Default value from credhub*) Exclude number characters from generated credential value.
- **include_special**: (*Optional*, Default: *Default value from credhub*) Include special characters from generated credential value.

---

### Generic

This generic resource has been made to prevent any future update on credhub which is could not yet implemented in this provider.

We don't recommend to use it if you can't secure your configuration, you will need to set credentials directly to your configuration.

Due to limitation to put arbitrary data inside a terraform parameter, there is 3 formats to create credentials in this resource:

To see what you can do on credhub you can have look at: http://credhub-api.cfapps.io

**Data credential format**
```hcl
resource "credhub_generic" "myuser" {
  type = "user"
  name = "/test/myuser"
  data_credential = {
    "username" = "FQnwWoxgSrDuqDLmeLpU"
    "password" = "6mRPZB3bAfb8lRpacnXsHfDhlPqFcjH2h9YDvLpL"
  }
}
```

**Data value format**
```hcl
resource "credhub_generic" "myvalue" {
  type = "value"
  name = "/test/myvalue"
  data_value = "myvalue"
}
```

**Data json format**
```hcl
resource "credhub_generic" "myjson" {
  type = "json"
  name = "/test/myjson"
  data_json = "{\"key\": {\"elem1\": \"value\", \"elem2\": \"value2\"}}"
}
```

- **name**: (**Required**) Name of your credential
- **type**: (**Required**) Type of your credential (see: http://credhub-api.cfapps.io/#set-credentials )
- **data_value**: (*Optional*, Default: *NULL*) A simple value as credential parameter. This can't be use with `data_json` or `data_credential`
- **data_credential**: (*Optional*, Default: *NULL*) A map with string values as credential parameter. This can't be use with `data_json` or `data_value`.
- **data_json**: (*Optional*, Default: *NULL*) A json string as credential parameter. This can't be use with `data_credential` or `data_value`.

---

## Datasources

**Note**: Computed parameters is what has been filled by the data source, this is what you can use after.

### Value

```hcl
data "credhub_value" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `value`.
- **value**: (*Computed*) Data value from your credential.

---

### Json

```hcl
data "credhub_json" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `json`.
- **json**: (*Computed*) Data in json format (as plain text) from your credential.

---

### Password

```hcl
data "credhub_password" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `password`.
- **password**: (*Computed*) Password from your credential.

---

### Certificate

```hcl
data "credhub_certificate" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `certificate`.
- **ca**: (*Computed*) CA from your credential.
- **ca_name**: (*Computed*) CA Name from your credential.
- **certificate**: (*Computed*) Certificate in pem format from your credential.
- **private_key**: (*Computed*) Private key in pem format from your credential.

---

### RSA

```hcl
data "credhub_rsa" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `rsa`.
- **public_key**: (*Computed*) Public key from your credential.
- **private_key**: (*Computed*) Private key from your credential.

---

### SSH

```hcl
data "credhub_ssh" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `ssh`.
- **public_key**: (*Computed*) Public key from your credential.
- **private_key**: (*Computed*) Private key from your credential.

---

### User

```hcl
data "credhub_user" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential.
- **cred_id**: (**Required if name empty**) Id of your credential.
- **type**: (*Computed*) This is the type of your credential but here it will be always `user`.
- **username**: (*Computed*) Username from your credential.
- **password**: (*Computed*) Password from your credential.

---

### Generic

This generic datasource has been made to support any future type on credhub which is could not yet implemented in this provider.

To see what you can do on credhub you can have look at: http://credhub-api.cfapps.io

```hcl
data "credhub_generic" "my_data" {
  name = "mydata"
  // or you can use credential id:
  // cred_id = "mydata-id"
}
```

- **name**: (**Required if cred_id empty**) Name of your credential
- **cred_id**: (**Required if name empty**) Id of your credential
- **type**: (*Computed*) This is the type of your credential
- **credential**: (*computed*) Data retrieve from your crendential. This is a map of string, a type value will be accessible by `credential.value`, data with a subtree will be converted to json if needed.
