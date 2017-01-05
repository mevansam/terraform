---
layout: "bosh"
page_title: "Bosh: bosh_vip_network"
sidebar_current: "docs-bosh-resource-infrastructure-vip-network"
description: |-
  Provides a Bosh VIP Network resource.
---

# bosh\_vip\_network

Provides a Bosh [VIP Network](http://bosh.io/docs/networks.html#vip) resource. You use a VIP network to map public IPs such as OpenStack Neutron floating IPs or AWS elastic IPs to instances.

## Example Usage

The following declares public network.

```
# Configure a VIP network to map public IPs to instance

resource "bosh_vip_network" "public" {    
    name = "public"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the network

IaaS specific arguments:
> Multiple IaaS specific attributes can be declared and the correct one will be selected based on the target's CPI specification.

* `azure` - (Optional) Azure specific attributes

  - `resource_group_name` - (String, optional) Name of a resource group. If it is set, Azure CPI will search the public IP in this resource group. Otherwise, Azure CPI will search the public IP in resource_group_name in the global CPI settings.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the Network
