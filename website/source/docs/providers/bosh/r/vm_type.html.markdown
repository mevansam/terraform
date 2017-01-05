---
layout: "bosh"
page_title: "Bosh: bosh_vm_type"
sidebar_current: "docs-bosh-resource-infrastructure-vm-type"
description: |-
  Provides a Bosh VM Type resource.
---

# bosh\_vm\_type

Provides a Bosh [VM Type](http://bosh.io/docs/cloud-config.html#vm-types) resource.

## Example Usage

The following declares a large VM Type for AWS and OpenStack.

```
# Configure the Bosh Large VM Type

resource "bosh_vm_type" "large" {
    
    name = "large"
    
    aws {
        instance_type = "m4.large"
    }
    openstack {
        instance_type = "m2.large"
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name identifying the VM Type

IaaS specific arguments per `subnet` or at the resource level:
> Multiple IaaS specific arguments can be declared and the correct one will be selected based on the target's CPI specification.

* `aws` - (Optional) AWS specific attributes

  * `instance_type` - (String, required) Type of the instance. Example: m3.medium.

  * `availability_zone` - (String, optional) Availability zone to use for creating instances. Example: us-east-1a.

  * `security_groups` - (Array, optional) See description under networks. Available in v46+.

  * `key_name` - (String, optional) Key pair name. Defaults to key pair name specified by default_key_name in global CPI settings. Example: bosh.

  * `spot_bid_price` - (Float, optional) Bid price in dollars for AWS spot instance. Using this option will slow down VM creation. Example: 0.03.

  * `spot_ondemand_fallback` - (Boolean, optional) Set to true to use an on demand instance if a spot instance is not available during VM creation. Defaults to false. Available in v36.

  * `elbs` - (Array, optional) Array of ELB names that should be attached to created VMs.

  * `iam_instance_profile` - (String, optional) Name of an IAM instance profile. Example: director.

  * `placement_group` - (String, optional) Name of a placement group. Example: my-group.

  * `tenancy` - (String, optional) VM tenancy configuration. Example: dedicated. Default is default.

  * `auto_assign_public_ip` - (Boolean, optional) Assigns a public IP address to the created VM. This IP is ephemeral and may change; use an Elastic IP instead for a persistent address. Defaults to false. Available in v55+.

  * `advertised_route` - (Array, optional) Creates routes in an AWS Route Table with the created BOSH VM as the target.

      - `table_id` - (String, required) ID of the route table in which to create the route (e.g. rt-abcdef123).
      - `destination` - (String, required) Destination CIDR for the route. All traffic with a destination within this CIDR will be routed through the created BOSH VM.

  * `raw_instance_storage` - (Boolean, optional) Exposes all available instance storage via labeled disks. Defaults to false.

  * `source_dest_check` - (Boolean, optional) Specifies whether the instance must be the source or destination of any traffic it sends or receives. If set to false, the instance does not need to be the source or destination. Used for network address translation (NAT) boxes, frequently to communicate between VPCs. Defaults to true. Available in v59+.

  * `ephemeral_disk` - (Optional) EBS backed ephemeral disk of custom size. Default disk size is either the size of first instance storage disk, if the instance_type offers it, or 10GB. Before v53: Used EBS only if instance storage is not large enough or not available for selected instance type.

      - `size` - (Integer, required) Specifies the disk size in megabytes.
      - `type` - (String, optional) Type of the disk: standard, gp2. Defaults to gp2.
      - `iops` - (Integer, optional) Specifies the number of I/O operations per second to provision for the drive.
      - `encrypted` - (Boolean, optional) Enables encryption for the EBS backed ephemeral disk. An error is raised, if the instance_type does not support it. Since v53. Defaults to false.
      - `use_instance_storage` - (Boolean, optional) Forces the usage of instance storage as ephemeral disk backing. Will raise an error, if the used instance_type does not have instance storage. Cannot be combined with any other option under ephemeral_disk or with raw_instance_storage. Since v53. Defaults to false.

  * `root_disk` - (Optional) EBS backed root disk of custom size.

      - `size` - (Integer, required) Specifies the disk size in megabytes.
      - `type` - (String, optional) Type of the disk: standard, gp2. Defaults to gp2.

* `azure` - (Optional) Azure specific attributes

  * `instance_type` - (String, required) Type of the instance. Example: Standard_A2.
  
  * `storage_account_name` - (String, optional) Storage account for VMs. If this is not set, the VMs will be created in the default storage account. See this document for more details on why this option exists.

  * `storage_account_max_disk_number` - (Integer, optional) Number of disks limitation in a storage account. Default value is 30. This will be used only when storage_account_name is a pattern.

  * `storage_account_type` - (String, optional) Storage account type. It is required if the storage account does not exist. It can be either Standard_LRS, Standard_ZRS, Standard_GRS, Standard_RAGRS or Premium_LRS.

  * `storage_account_location` - (String, optional) Location of the storage account. It is needed if the storage account does not exist. If it is not set, the location of the default resource group will be used. 

  * `availability_set` - (String, optional) Name of an availability set to use for VMs. 

  * `platform_update_domain_count` - (Integer, optional) The count of update domain in the availability set. Default value is 5.

  * `platform_fault_domain_count` - (Integer, optional) The count of fault domain in the availability set. Default value is 3.

  * `load_balancer` - (String, optional) Name of an load balancer the VMs should belong to. You need to create the load balancer manually before configuring it.

  * `security_group` - (String, optional]) The security group to apply to network interfaces of all VMs placed in this resource pool. The security group of a network interface can be specified either in a resource pool(higher priority) or a network configuration(lower priority); if it is not specified, the default security group (specified by default_security_group in the global CPI settings) will be used.

  * `caching` - (String, optional) Type of the disk caching of the VMs’ OS disks. It can be either None, ReadOnly or ReadWrite. Default is ReadWrite.

  * `root_disk` - (Optional) OS disk of custom size.

      - `size` - (Integer, optional]:) Specifies the disk size in MiB.

  * `ephemeral_disk` - (Optional) Ephemeral disk to apply for all VMs that are in this resource pool. By default a data disk with the default size as below will be created as the ephemeral disk.

      - `use_root_disk` - (Boolean, optional) Enable to use OS disk to store the ephemeral data. The default value is false. When it is true, ephemeral_disk.size will not be used.
      - `size` - (Integer, optional) Specifies the disk size in MiB. If this is not set, the default size as below will be used. The size of the ephemeral disk for the BOSH VM should be larger than or equal to 30*1024 MiB. Please always use N * 1024 as the size because Azure always uses GiB not MiB.

  * `assign_dynamic_public_ip` - (Boolean, optional) Enable to create and assign dynamic public IP to the VM automatically (to solve the azure SNAT issue). Default value is false. Only the VM without vip will be assigned a dynamic public IP when this value is set to true, and the dynamic public IP will be deleted when the VM is deleted.

* `openstack` - (Optional) OpenStack specific attributes

  * `instance_type` - (String, required) Type of the instance. Example: m1.small.

  * `availability_zone` - (String, optional) Availability zone to use for creating instances. Example: east.

  * `security_groups` - (Array, optional) Array of security groups to apply for all VMs that are in this resource pool. Defaults to security groups specified by default_security_groups in the global CPI settings unless security groups are specified on one of the VM networks. Security groups can be specified either on a resource pool or on a network. Available in v16+.

  * `key_name` - (String, optional) Key pair name. Defaults to key pair name specified by default_key_name in the global CPI settings. Example: bosh.

  * `scheduler_hints` - (Hash, optional) Data passed to the OpenStack Filter scheduler to influence its decision where new VMs can be placed. See VM Anti-Affinity for a detailed example. Example: { group: af09abf2-2283... }

  * `root_disk` - (Hash, optional) Custom root disk properties. Requires boot_from_volume: true to enable cinder-backed boot volumes. Available in v25+.

      - `size` - (Integer, required) Specifies the disk size in gigabytes.

* `vcenter` - (Optional) vCenter specific attributes

  * `cpu` - (Integer, required) Number of CPUs. Example: 1.

  * `ram` - (Integer, required) RAM in megabytes. Example: 1024.

  * `disk` - (Integer, required) Ephemeral disk size in megabytes. Example: 10240.

  * `cpu_hot_add_enabled` - (Boolean, optional) Allows operator to add additional CPU resources while the VM is on. Default: false. Available in v21+.

  * `memory_hot_add_enabled` - (Boolean, optional) Allows operator to add additional memory resources while the VM is on. Default: false. Available in v21+.

  * `nested_hardware_virtualization` - (Boolean, optional) Exposes hardware assisted virtualization to the VM. Default: false.

  * `datastores` - (Array, optional) Allows operator to specify a list of ephemeral datastores for the VM. These names are exact datastore names and not regex 
  patterns. At least one of these datastores must be accessible from clusters provided in resource_pools.cloud_properties/azs.cloud_properties or in the global CPI configuration. Available in v23+.

  * `datacenters` - (Array, optional) Used to override the VM placement specified under azs.cloud_properties. The format is the same as under AZs.

  * `nsx` - (Optional) VMWare NSX additions section. Available in CPI v30+ and NSX v6.1+.

      - `security_groups` - (Array, optional) A collection of security group names that the instances should belong to. The CPI will create the security groups if they do not exist. BOSH will also automatically create security groups based on metadata such as deployment name and instance group name. The full list of groups can be seen under create_vm’s environment groups.
      - `lb` - (Array, optional) A collection of NSX Edge Load Balancers (LBs) to which instances should be attached. The LB and Server Pool must exist prior to the deployment.

          - `edge_name` - (String, required) Name of the NSX Edge.
          - `pool_name` - (String, required) Name of the Edge’s Server Pool.
          - `security_group` - (String, required) Name of the Pool’s target Security Group. The CPI will add the VM to the specified security group (creating the security group if needed), then add the security group to the specified Server Pool.
          - `port` - (Integer, required) The port that the VM’s service is listening on (e.g. 80 for HTTP).
          - `monitor_port` - (Integer, optional) The healthcheck port that the VM is listening on. Defaults to the value of port.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the VM Type