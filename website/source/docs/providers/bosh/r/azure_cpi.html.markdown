---
layout: "bosh"
page_title: "Bosh: bosh_azure_cpi"
sidebar_current: "docs-bosh-resource-cpi-azure"
description: |-
  Provides a Bosh Azure Cloud Provider Interface resource.
---

# bosh\_azure\_cpi

Provides a Bosh Azure [Cloud Provider Interface](http://bosh.io/docs) resource which is an abstraction of the underlying IaaS. It can be used to manage Bosh deployments on a Microsoft Azure IaaS environment.

## Example Usage

```
resource "bosh_azure_cpi" "azure" {    
    
    ...
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (required) e.g. '3c39a033-c306-4615-a4cb-260418d63879'
* `tenant_id` - (required) e.g. '0412d4fa-43d2-414b-b392-25d5ca46561da'
* `client_id` - (required) e.g. '33e56099-0bde-8z93-a005-89c0f6df7465'
* `client_secret` - (required) e.g. 'client-secret'
* `resource_group_name` - (required) e.g. 'bosh-res-group'
* `storage_account_name` - (required) e.g. 'boshstore'
* `default_security_group` - (required) e.g. '+o9EVnno8ja5OzQim6fQkVGbvoQmazl+Nlg=='
* `vnet_name` - (optional) Name of created Virtual Network
* `subnet_name` - (optional) Deployment manifest assumes that the subnet is 10.0.0.0/24 and Director VM will be placed at 10.0.0.4
* `public_ip` - (optional) The IP address of created public IP for BOSH
* `ssh_user` - (required) The ssh user (e.g. vcap)
* `ssh_public_key` - (required) The generated SSH public key

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the CPI
