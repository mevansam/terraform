---
layout: "bosh"
page_title: "Bosh: bosh_disk_type"
sidebar_current: "docs-bosh-resource-infrastructure-disk-type"
description: |-
  Provides a Bosh Disk Type resource.
---

# bosh\_disk\_type

Provides a Bosh [Disk Type](http://bosh.io/docs/cloud-config.html#disk-types) resource.

## Example Usage

The following declares a 20GB disk type that is used by Bosh to create persistant disks.

```
resource "bosh_disk_type" "small" {
    
    name = "small"
    disk_size = 20480
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name identifying the Disk Type
* `disk_size` - (Integer, required) Specifies the disk size. disk_size must be a positive integer. BOSH creates a persistent disk of that size in megabytes and attaches it to each job instance VM.

IaaS specific arguments per `subnet` or at the resource level:
> Multiple IaaS specific arguments can be declared and the correct one will be selected based on the target's CPI specification.

* `aws` - (Optional) AWS specific attributes

  * `type` - (String, optional) Type of the disk: standard, gp2. Defaults to gp2.
  * `iops` - (Integer, optional) Specifies the number of I/O operations per second to provision for the drive.
  * `encrypted` - (Boolean, optional) Turns on EBS volume encryption for this persistent disk. VM root and ephemeral disk are not encrypted. Defaults to false.
  * `kms_key_arn` - (String, optional) Encrypts the disk using an encryption key stored in the AWS Key Management Service (KMS). The format of the ID is XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX. Be sure to use the Key ID, not the Alias. If this property is omitted and encrypted is true, the disk will be encrypted using your account’s default aws/ebs encryption key.

* `azure` - (Optional) Azure specific attributes

  * `caching` - (String, optional) Type of the disk caching. It can be either None, ReadOnly or ReadWrite. Default is None. Only None and ReadOnly are supported for premium disks.

* `openstack` - (Optional) OpenStack specific attributes

  * `type` - (String, optional) Volume type as configured in your OpenStack installation. Example: SSD

* `vcenter` - (Optional) vCenter specific attributes

  * `type` - (String, optional) Type of the disk: thick, thin, preallocated, eagerZeroedThick. Defaults to preallocated. Available in v12. Overrides the global default_disk_type.

  * `datastores` - (Array, optional) List of datastore names for storing persistent disks. Overrides the global persistent_datastore_pattern. These names are exact datastore names and not regex patterns. Available in v29+.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the Disk Type
