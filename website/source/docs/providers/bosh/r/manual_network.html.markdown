---
layout: "bosh"
page_title: "Bosh: bosh_manual_network"
sidebar_current: "docs-bosh-resource-infrastructure-manual-network"
description: |-
  Provides a Bosh Manual Network resource.
---

# bosh\_manual\_network

Provides a Bosh [Manual Network](http://bosh.io/docs/networks.html#manual) resource. In a manual network IPs are managed by Bosh using on the provided network parameters.

## Example Usage

The following declares private manual network for a deployment on OpenStack.

```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "subnet_cidr" {}
variable "subnet_gateway" {}
variable "subnet_dns" {}

variable "openstack_network_uuid" {}
variable "openstack_bosh_security_group" {}

# Configure a manual network for a private IP network

resource "bosh_manual_network" "private" {

    name = "private"

    subnet {
        range = "${var.subnet_cidr}"
        gateway = "${var.subnet_gateway}"
        dns = [ "${var.subnet_dns}" ]

        openstack {
            net_id = "${var.openstack_network_uuid}"
            security_groups: [ "${var.openstack_bosh_security_group}" ]
        }
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the network

* `subnet` - (Array, required) Lists subnets in this network

  - `range` - (String, optional/computed) Subnet IP range that includes all IPs from this subnet
  - `gateway` - (String, optional/computed) Subnet gateway IP
  - `dns` - (Array, optional) DNS IP addresses for this subnet
  - `reserved` - (Array, optional) Array of reserved IPs and/or IP ranges. BOSH does not assign IPs from this range to any VM
  - `static` - (Array, optional) Array of static IPs and/or IP ranges. BOSH assigns IPs from this range to jobs requesting static IPs. Only IPs specified here can be used for static IP reservations.
  - `azs` - (Array, optional/computed) List of AZs associated with this subnet (should only be used when using first class AZs). Example: [z1, z2]. Available in v241+.

IaaS specific arguments for each `subnet`:
> Multiple IaaS specific arguments can be declared and the correct one will be selected based on the target's CPI specification.

* `aws` - (Optional) AWS specific attributes

  - `subnet` - (String, required) Subnet ID in which the instance will be created. Example: subnet-9be6c3f7.
  - `security_groups` - (Array, optional) Array of Security Groups, by name or ID, to apply to all VMs placed on this network. Security groups can be specified as follows, ordered by greatest precedence: vm_types, followed by networks, followed by default_security_groups.

* `azure` - (Optional) Azure specific attributes

  - `resource_group_name` - (String, optional) Name of a resource group. If it is set, Azure CPI will search the virtual network and security group in this resource group. Otherwise, Azure CPI will search the virtual network and security group in resource_group_name in the global CPI settings.
  - `virtual_network_name` - (String, required) Name of a virtual network. Example: boshnet.
  - `subnet_name` - (String, required) Name of a subnet within virtual network.
  - `security_group` - (String, optional) The security group to apply to network interfaces of all VMs placed in this network. The security group of a network interface can be specified either in a resource pool(higher priority) or a network configuration(lower priority); if it is not specified, the default security group (specified by default_security_group in the global CPI settings) will be used.

* `openstack` - (Optional) OpenStack specific attributes

  - `net_id` - (String, required) Network ID containing the subnet in which the instance will be created. Example: net-b98ab66e-6fae-4c6a-81af-566e630d21d1.
  - `security_groups` - (Array, optional) Array of security groups to apply for all VMs that are placed on this network. Defaults to security groups specified by   - default_security_groups in the global CPI settings unless security groups are specified on a resource pool for a VM. Security groups can be specified either on a resource pool or on a network.

* `vcenter` - (Optional) vCenter specific attributes

  - `name` - (String, required) Name of the vSphere network. Example: 'VM Network'.

## Attributes Reference

The following attributes are exported:

* `id` -  A GUID identifying the Network