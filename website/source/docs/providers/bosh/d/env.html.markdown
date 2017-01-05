---
layout: "bosh"
page_title: "Bosh: bosh_env"
sidebar_current: "docs-bosh-datasource-env"
description: |-
  Targets a Bosh Director environment
---

# bosh\_env

The 'bosh_env' data source is used to target an existing Director.

## Example Usage

```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "bosh_user" {}
variable "bosh_password" {}

# Configure the Bosh Env Target

data "bosh_env" "e1" {
    target = "${bosh_deployment.bosh.instance_ips["bosh"]}"
    user = "${var.bosh_user}"
    password = "${var.bosh_password}"
}
```

## Argument Reference

The following arguments are supported:

* `target` - (required) Bosh target. This can also be specified with the `BOSH_TARGET` shell environment variable. This target can point to an endpoint which has been deployed by a [`bosh_deployment`](/docs/providers/bosh/r/deployment.html) resource, as shown in the example above. By referencing the deployment resource you ensure that the bosh environment has deployed correctly before using it manage distributed deployments.

* `user` - (optional) The Bosh user. Defaults to "admin". This can also be 
specified with the `BOSH_USER` shell environment variable.

* `password` - (optional) The Bosh user's password. This can also be specified 
with the `BOSH_PASSWORD` shell environment variable.

## Attributes Reference

The following attributes are exported:

* `name` - Name of the Director.
* `uuid` - Unique ID of the Director.
* `version` - Version of the Director software.
* `cpi` -  Name of the CPI the Director will use.
* `user_authentication_type` - Type of the authentication the Director is configured to expect.
* `user_authentication_options` - Map of name-value pairs providing additional information to 
how authentication should be performed.
