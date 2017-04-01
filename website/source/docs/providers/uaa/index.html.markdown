---
layout: "uaa"
page_title: "Provider: UAA"
sidebar_current: "docs-uaa-index"
description: |-
  The UAA provider is used to manage client and user accounts on a User Account and Authentication (UAA) identitity manageement service. The provider needs to be configured with the proper credentials before it can be used.
---

# User Account and Authentication (UAA) Provider

The UAA provider is used to interact with a [User Account and Authentication (UAA)](https://github.com/cloudfoundry/uaa)  
identitity manageement service. The resources associated with this provider can be used to configure OAuth clients and
users managed by the service. They are based off of the published [UAA API Spec](https://docs.cloudfoundry.org/api/uaa/).

The UAA provider is typically used along with the [CF provider](/docs/providers/cf/index.html) to configure a Cloud Foundry 
installation.

Use the navigation to the left to read about the available resources.

## Example Usage

```
# Set the variable values in *.tfvars file
# or using -var="api_url=..." CLI option

variable "uaa_api_url" {}
variable "uaa_admin_client_secret" {}

# Configure the CloudFoundry Provider

provider "cloudfoundry" {
    uaa_api_url = "${var.api_url}"
    client_id = "admin"
    client_secret = "${var.admin_password}"
    skip_ssl_validation = true
}
```

## Argument Reference

The following arguments are supported:

* `uaa_api_url` - (Required) API endpoint (e.g. https://uaa.local.pcfdev.io). This can also be specified
  with the `UAA_API_URL` shell environment variable.

* `client_id` - (Optional) The UAA admin client ID. Defaults to "admin". This can also be specified
  with the `UAA_CLIENT_ID` shell environment variable.

* `client_secret` - (Required) The secret of the UAA admin client. This can also be specified
  with the `UAA_CLIENT_SECRET` shell environment variable.

* `skip_ssl_validation` - (Optional) Skip verification of the API endpoint - Not recommended!. Defaults to "false". This can also be specified
  with the `UAA_SKIP_SSL_VALIDATION` shell environment variable.
