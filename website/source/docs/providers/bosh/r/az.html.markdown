---
layout: "bosh"
page_title: "Bosh: bosh_az"
sidebar_current: "docs-bosh-resource-infrastructure-az"
description: |-
  Provides a Bosh Availability Zone resource.
---

# bosh\_az

Provides a Bosh [Availability Zone](http://bosh.io/docs/cloud-config.html#azs) resource.

## Example Usage

The following declares a Bosh AZ in OpenStack.

```
resource "bosh_az" "z1" {
    
    name = "z1"

    openstack {
        availability_zone = "AZ1"
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Availability Zone (AZ)

IaaS specific arguments:
> Multiple IaaS specific arguments can be declared and the correct one will be selected based on the target's CPI specification.

* `aws` - (Optional) Attributes that maps this Bosh AZ specification to a physical AWS AZ.

  - `availability_zone` - (String, required) Availability zone to use for creating instances. e.g. us-east-1a

* `azure` - (N/A) Currently the CPI does not support any cloud properties for AZs.

* `openstack` - (Optional) Attributes that maps this Bosh AZ specification to a physical OpenStack AZ.

  - `availability_zone` - (String, required) Availability zone to use for creating instances. e.g. AZ1

* `vcenter` - (Optional) Attributes that maps this Bosh AZ specification to a physical vcenter Datacenter cluster.

  - `name` (String, required) vSphere datacenter name.
  - `cluster` (Array, required) vSphere host cluster to use as this AZ
  
      - `name` - (String, required) Name of cluster in vSphere
      - `resource_pool` - (String, optional) Name of vSphere Resource Pool to use for VM placement.
      - `drs_rules` - (Array, optional) Array of DRS rules applied to constrain VM placement. Must have only one.

          - `name` - (String, required) Name of a DRS rule that the Director will create.
          - `type` - (String, required) Type of a DRS rule. Currently only a value of 'separate_vms' is supported.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the AZ
